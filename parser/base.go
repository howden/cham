package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
	"strconv"
)

// Parses a variable - either a number or an identifier
func (parser *Parser) parseVariable() (ast.IntegerTerm, error) {
	if parser.currentToken.Type == token.Number {
		i, err := parser.parseNumber()
		if err != nil {
			return nil, err
		}
		return ast.Number(i), nil
	}

	if parser.currentToken.Type == token.Ident {
		ident, err := parser.parseIdent()
		if err != nil {
			return nil, err
		}
		return ast.Ident(ident), nil
	}

	return nil, fmt.Errorf("expected number or ident but got %v instead", parser.currentToken)
}

func (parser *Parser) parseNumber() (int, error) {
	sign := 1
	if parser.currentToken.Type == token.Subtract {
		parser.next()
		sign = -1
	}

	ok, err := parser.expectToken(token.Number)
	if !ok {
		return 0, err
	}

	i, err := strconv.Atoi(parser.currentToken.Literal)
	if err != nil {
		return 0, fmt.Errorf("parser error for int: %v, %w", parser.currentToken.Literal, err)
	}

	parser.next()
	return sign * i, nil
}

func (parser *Parser) parseIdent() (string, error) {
	ok, err := parser.expectToken(token.Ident)
	if !ok {
		return "", err
	}

	ident := parser.currentToken.Literal
	parser.next()
	return ident, nil
}
