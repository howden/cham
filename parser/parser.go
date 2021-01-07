package parser

import (
	"fmt"
	"github.com/howden/cham/lexer"
	"github.com/howden/cham/token"
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
		return false, fmt.Errorf("expected token %v but got %v instead", expected, token.Type)
	}
	return true, nil
}

// Tests whether the current token matches the expected token type
func (parser *Parser) expectToken(expected token.TokenType) (ok bool, err error) {
	return expect(parser.currentToken, expected)
}
