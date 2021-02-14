package parser

import (
	"bytes"
	"fmt"
	"github.com/howden/cham/lexer"
	"github.com/howden/cham/token"
	"github.com/pkg/errors"
	"strings"
)

type Parser struct {
	lexer *lexer.Lexer

	// Holds the current token under consideration by the parser
	currentToken token.Token
}

// Creates a new parser using the given Lexer as a source of input tokens
func NewParser(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.next()
	return p
}

// Advance to the next token
func (parser *Parser) next() {
	parser.currentToken = parser.lexer.NextToken()
}

// Tests whether the token matches the expected token type
func expect(token token.Token, expected token.TokenType) (ok bool, err error) {
	if token.Type != expected {
		return false, fmt.Errorf("expected token %v but got %v instead", expected, token)
	}
	return true, nil
}

// Tests whether the current token matches the expected token type
func (parser *Parser) expectToken(expected token.TokenType) (ok bool, err error) {
	return expect(parser.currentToken, expected)
}

// Custom error wrapper which additionally contains information
// about the state of the lexer when the error occurred.
// (where in the src code caused the parsing error!)
type ParserError struct {
	cause              error
	LexerCurrentLine   int
	LexerCurrentColumn int
}

func (err *ParserError) Error() string {
	return fmt.Sprintf("error while parsing (at %d:%d): %v", err.LexerCurrentLine, err.LexerCurrentColumn, err.cause)
}

func (err *ParserError) Cause() error {
	return err.cause
}

func (err *ParserError) Unwrap() error {
	return err.cause
}

// function to create an error wrapper given the parser
func (parser *Parser) wrapError(err error) *ParserError {
	err = errors.WithStack(err)
	line, col := 0, 0

	pos := parser.lexer.Pos()
	if pos.IsValid() {
		line = pos.Line
		col = pos.Column
	}

	return &ParserError{cause: err, LexerCurrentLine: line, LexerCurrentColumn: col}
}

func PrintParserError(src string, err error) {
	if pe, ok := err.(*ParserError); ok {
		fmt.Println(err)
		fmt.Printf("\n%s\n", src)
		fmt.Printf("%s^ HERE\n", strings.Repeat(" ", pe.LexerCurrentColumn-2))
	} else {
		fmt.Println(err)
	}
}

func FormatErrorWithParserLocation(err error) error {
	if pe, ok := err.(*ParserError); ok {
		var buf bytes.Buffer
		_, _ = fmt.Fprintf(&buf, "%s^ ", strings.Repeat(" ", pe.LexerCurrentColumn-1))
		_, _ = fmt.Fprint(&buf, err)
		return errors.New(buf.String())
	} else {
		return err
	}
}
