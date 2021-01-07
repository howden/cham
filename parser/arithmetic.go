package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
)

// <addop> ::= <plus> | <subtract>
// <multop> ::= <multiply> | <divide>
// <aexp> ::= <aterm> {<addop> <aexp>}
// <aterm> ::= <afactor> {<multop> <aterm>}
// <afactor> ::= <variable>
// <afactor> ::= <openb> <aexp> <closeb>

func isAddOp(typ token.TokenType) bool {
	return typ == token.Plus || typ == token.Subtract
}

func isMultOp(typ token.TokenType) bool {
	return typ == token.Multiply || typ == token.Divide
}

func (parser *Parser) parseAexp() (ast.IntegerTerm, error) {
	root, err := parser.aterm()
	if err != nil {
		return nil, err
	}

	for isAddOp(parser.currentToken.Type) {
		opType := parser.currentToken.Type
		parser.next()

		left := root
		right, err := parser.parseAexp()
		if err != nil {
			return nil, err
		}

		if opType == token.Plus {
			root = ast.Plus(left, right)
		} else {
			root = ast.Subtract(left, right)
		}
	}

	return root, nil
}

func (parser *Parser) aterm() (ast.IntegerTerm, error) {
	root, err := parser.afactor()
	if err != nil {
		return nil, err
	}

	for isMultOp(parser.currentToken.Type) {
		opType := parser.currentToken.Type
		parser.next()

		left := root
		right, err := parser.aterm()
		if err != nil {
			return nil, err
		}

		if opType == token.Multiply {
			root = ast.Multiply(left, right)
		} else {
			root = ast.Divide(left, right)
		}
	}

	return root, nil
}

func (parser *Parser) afactor() (ast.IntegerTerm, error) {
	if parser.currentToken.Type == token.OpenBracket {
		parser.next()
		exp, err := parser.parseAexp()
		if err != nil {
			return nil, err
		}
		if parser.currentToken.Type != token.CloseBracket {
			return nil, fmt.Errorf("expected close bracket but got %v instead", parser.currentToken)
		}
		parser.next()
		return exp, nil
	} else {
		variable, err := parser.parseVariable()
		if err != nil {
			return nil, err
		}

		return variable, nil
	}
}
