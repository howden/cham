package lexer

import (
	"fmt"
	"github.com/howden/cham/token"
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

type Lexer struct {
	scanner *scanner.Scanner
}

// Creates a new Lexer from an input string
func FromString(input string) *Lexer {
	return FromReader(strings.NewReader(input), "repl")
}

// Creates a new Lexer from an input io.Reader, with the given file name
func FromReader(input io.Reader, fileName string) *Lexer {
	var s scanner.Scanner
	s.Init(input)
	s.Filename = fileName
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | /*scanner.ScanChars | scanner.ScanStrings |*/ scanner.ScanComments | scanner.SkipComments
	s.IsIdentRune = func(ch rune, i int) bool {
		return unicode.IsLetter(ch) /*|| unicode.IsDigit(ch) && i > 0*/
	}

	l := &Lexer{&s}
	return l
}

// "Simple" tokens with a direct, non-ambiguous mapping from single character to token
var simpleTokens = map[rune]token.TokenType{
	'!': token.Not,
	'(': token.OpenBracket,
	')': token.CloseBracket,
	'{': token.OpenCurlyBracket,
	'}': token.CloseCurlyBracket,
	'+': token.Plus,
	'-': token.Subtract,
	'*': token.Multiply,
	'/': token.Divide,
	',': token.Comma,
}

// Produces the next token from the lexer
func (lexer *Lexer) NextToken() (token.Token, error) {
	s := lexer.scanner
	tok := s.Scan()

	if tok == scanner.EOF {
		return token.EOF.New(), nil
	} else if tok == scanner.Ident {
		if lexer.scanner.TokenText() == "if" {
			return token.If.New(), nil
		}
		return token.Ident.WithLiteral(s.TokenText()), nil
	} else if tok == scanner.Int {
		return token.Number.WithLiteral(s.TokenText()), nil
	} else if tok == '=' && s.Peek() == '>' {
		s.Scan()
		return token.ReactionOp.New(), nil
	} else if tok == '<' {
		if s.Peek() == '=' {
			s.Scan()
			return token.LessThanOrEqual.New(), nil
		}
		return token.LessThan.New(), nil
	} else if tok == '>' {
		if s.Peek() == '=' {
			s.Scan()
			return token.GreaterThanOrEqual.New(), nil
		}
		return token.GreaterThan.New(), nil
	} else if tok == '=' && s.Peek() == '=' {
		s.Scan()
		return token.Equal.New(), nil
	} else if tok == '!' && s.Peek() == '=' {
		s.Scan()
		return token.NotEqual.New(), nil
	} else if tok == '|' && s.Peek() == '|' {
		s.Scan()
		return token.Or.New(), nil
	} else if tok == '&' && s.Peek() == '&' {
		s.Scan()
		return token.And.New(), nil
	} else if desc, found := simpleTokens[tok]; found {
		return desc.New(), nil
	} else {
		return token.Token{}, fmt.Errorf("unknown token '%s' at %s", s.TokenText(), s.Pos())
	}
}

// Consumes the remaining tokens from the lexer until EOF is reached, and returns them in a slice
func (lexer *Lexer) RemainingTokens() ([]token.Token, error) {
	var arr []token.Token
	for tok, err := lexer.NextToken(); err != nil || tok.Type != token.EOF; tok, err = lexer.NextToken() {
		if err != nil {
			return nil, err
		}
		arr = append(arr, tok)
	}
	return arr, nil
}
