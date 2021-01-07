package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
)

var comparisonAsts = map[token.TokenType]func(left ast.IntegerTerm, right ast.IntegerTerm) ast.BooleanTerm{
	token.Equal:              ast.Equals,
	token.NotEqual:           ast.NotEquals,
	token.LessThan:           ast.LessThan,
	token.GreaterThan:        ast.GreaterThan,
	token.LessThanOrEqual:    ast.LessThanEqual,
	token.GreaterThanOrEqual: ast.GreaterThanEqual,
}

func (parser *Parser) parseComparison() (ast.BooleanTerm, error) {
	left, err := parser.parseAexp()
	if err != nil {
		return nil, err
	}

	fun, ok := comparisonAsts[parser.currentToken.Type]
	if !ok {
		return nil, fmt.Errorf("expected comparison operator but got %v", parser.currentToken)
	}

	right, err := parser.parseAexp()
	if err != nil {
		return nil, err
	}

	return fun(left, right), nil
}
