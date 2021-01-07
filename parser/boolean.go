package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
)

// <bexp> ::= <bterm> {<or> <bterm>}
// <bterm> ::= <bnotfactor> {<and> <bnotfactor>}
// <bnotfactor> ::= <not> <bfactor>
// <bnotfactor> ::= <bfactor>
// <bfactor> ::= <bool-value>
// <bfactor> ::= <openb> <bexp> <closeb>

func (parser *Parser) parseBexp() (ast.BooleanTerm, error) {
	root, err := parser.bterm()
	if err != nil {
		return nil, err
	}

	for parser.currentToken.Type == token.Or {
		parser.next()
		left := root
		right, err := parser.bterm()
		if err != nil {
			return nil, err
		}
		root = ast.BooleanOr(left, right)
	}

	return root, nil
}

func (parser *Parser) bterm() (ast.BooleanTerm, error) {
	root, err := parser.notFactor()
	if err != nil {
		return nil, err
	}

	for parser.currentToken.Type == token.And {
		parser.next()
		left := root
		right, err := parser.notFactor()
		if err != nil {
			return nil, err
		}
		root = ast.BooleanAnd(left, right)
	}

	return root, nil
}

func (parser *Parser) notFactor() (ast.BooleanTerm, error) {
	if parser.currentToken.Type != token.Not {
		return parser.bfactor()
	}

	parser.next()

	fac, err := parser.bfactor()
	if err != nil {
		return nil, err
	}

	return ast.BooleanNot(fac), nil
}

func (parser *Parser) bfactor() (ast.BooleanTerm, error) {
	if parser.currentToken.Type == token.OpenBracket {
		parser.next()
		exp, err := parser.parseBexp()
		if err != nil {
			return nil, err
		}
		if parser.currentToken.Type != token.CloseBracket {
			return nil, fmt.Errorf("expected close bracket but got %v instead", parser.currentToken)
		}
		parser.next()
		return exp, nil
	} else {
		comp, err := parser.parseComparison()
		if err != nil {
			return nil, err
		}

		return comp, nil
	}
}
