package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
	"strconv"
)

// Parses a variable - either a number or an identifier
func (parser *Parser) parseVariable() (ast.IntegerTerm, error) {
	if ok, _ := parser.expectToken(token.Number); ok {
		i, err := strconv.Atoi(parser.currentToken.Literal)
		if err != nil {
			return nil, fmt.Errorf("parser error for int: %v, %w", parser.currentToken.Literal, err)
		}

		parser.next()
		return ast.Number(i), nil
	}

	if ok, _ := parser.expectToken(token.Ident); ok {
		ident := ast.Ident(parser.currentToken.Literal)
		parser.next()
		return ident, nil
	}

	return nil, fmt.Errorf("expected number or ident but got %v instead", parser.currentToken)
}
