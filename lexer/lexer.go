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
		return (unicode.IsLetter(ch) && unicode.IsLower(ch)) || ch == '_' /*|| unicode.IsDigit(ch) && i > 0*/
	}

	return &Lexer{&s}
}

// "Simple" tokens with a direct, non-ambiguous mapping from single character to token
var simpleTokens = map[rune]token.TokenType{
	'|': token.ReactionChain,
	':': token.ReactionDef,
	'!': token.Not,
	'(': token.OpenBracket,
	')': token.CloseBracket,
	'{': token.OpenCurlyBracket,
	'}': token.CloseCurlyBracket,
	'[': token.OpenSquareBracket,
	']': token.CloseSquareBracket,
	'+': token.Plus,
	'-': token.Subtract,
	'*': token.Multiply,
	'/': token.Divide,
	'%': token.Modulo,
	',': token.Comma,
}

// Produces the next token from the lexer
func (lexer *Lexer) NextToken() token.Token {
	s := lexer.scanner
	tok := s.Scan()

	if tok == scanner.EOF {
		return token.EOF.New()
	} else if tok == scanner.Ident {
		if lexer.scanner.TokenText() == "if" {
			return token.If.New()
		}
		return token.Ident.WithLiteral(s.TokenText())
	} else if tok == scanner.Int {
		return token.Number.WithLiteral(s.TokenText())
	} else if tok == '=' && s.Peek() == '>' {
		s.Scan()
		return token.ReactionOp.New()
	} else if tok == '<' {
		if s.Peek() == '=' {
			s.Scan()
			return token.LessThanOrEqual.New()
		}
		return token.LessThan.New()
	} else if tok == '>' {
		if s.Peek() == '=' {
			s.Scan()
			return token.GreaterThanOrEqual.New()
		}
		return token.GreaterThan.New()
	} else if tok == '=' && s.Peek() == '=' {
		s.Scan()
		return token.Equal.New()
	} else if tok == '!' && s.Peek() == '=' {
		s.Scan()
		return token.NotEqual.New()
	} else if tok == '|' && s.Peek() == '|' {
		s.Scan()
		return token.Or.New()
	} else if tok == '&' && s.Peek() == '&' {
		s.Scan()
		return token.And.New()
	} else if desc, found := simpleTokens[tok]; found {
		return desc.New()
	} else {
		return token.Error(fmt.Errorf("unknown token '%s' type %s at %s", s.TokenText(), scanner.TokenString(tok), s.Pos()))
	}
}

// Returns the current position of the scanner
func (lexer *Lexer) Pos() scanner.Position {
	return lexer.scanner.Pos()
}

// Consumes the remaining tokens from the lexer until EOF is reached, and returns them in a slice
func (lexer *Lexer) RemainingTokens() ([]token.Token, error) {
	var arr []token.Token
	for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
		if tok.Type == token.Invalid {
			return nil, tok.Err
		}
		arr = append(arr, tok)
	}
	return arr, nil
}
