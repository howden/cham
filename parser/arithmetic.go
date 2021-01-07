package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
)

// arithmetic.go contains the parsing code for arithmetic expressions.
//
// Context-free grammar accepted by this parser:
//   <aexp> ::= <aterm> {<addop> <aexp>}
//   <aterm> ::= <afactor> {<multop> <aterm>}
//   <afactor> ::= <variable>
//   <afactor> ::= <openb> <aexp> <closeb>

// Gets whether a token is an "addop"
func isAddOp(typ token.TokenType) bool {
	return typ == token.Plus || typ == token.Subtract
}

// Gets whether a token is an "multop"
func isMultOp(typ token.TokenType) bool {
	return typ == token.Multiply || typ == token.Divide
}

// Parses an "aexp"
// <aexp> ::= <aterm> {<addop> <aexp>}
func (parser *Parser) parseAexp() (ast.IntegerTerm, error) {
	root, err := parser.aterm()
	if err != nil {
		return nil, err
	}

	for isAddOp(parser.currentToken.Type) {
		// record the operator type, then advance the parser
		opType := parser.currentToken.Type
		parser.next()

		// use root as the left value, parse right from the next token
		left := root
		right, err := parser.parseAexp()

		if err != nil {
			return nil, err
		}

		// set root to the result
		if opType == token.Plus {
			root = ast.Plus(left, right)
		} else {
			root = ast.Subtract(left, right)
		}
	}

	return root, nil
}

// Parses an "aterm"
// <aterm> ::= <afactor> {<multop> <aterm>}
func (parser *Parser) aterm() (ast.IntegerTerm, error) {
	root, err := parser.afactor()
	if err != nil {
		return nil, err
	}

	for isMultOp(parser.currentToken.Type) {
		// record the operator type, then advance the parser
		opType := parser.currentToken.Type
		parser.next()

		// use root as the left value, parse right from the next token
		left := root
		right, err := parser.aterm()

		if err != nil {
			return nil, err
		}

		// set root to the result
		if opType == token.Multiply {
			root = ast.Multiply(left, right)
		} else {
			root = ast.Divide(left, right)
		}
	}

	return root, nil
}

// Parses an "afactor"
// <afactor> ::= <variable>
// <afactor> ::= <openb> <aexp> <closeb>
func (parser *Parser) afactor() (ast.IntegerTerm, error) {
	// try to parse brackets first
	if parser.currentToken.Type == token.OpenBracket {
		parser.next()

		// parse inner aexp
		exp, err := parser.parseAexp()
		if err != nil {
			return nil, err
		}

		// ensure bracket is closed after aexp is finished
		if parser.currentToken.Type != token.CloseBracket {
			return nil, fmt.Errorf("expected close bracket but got %v instead", parser.currentToken)
		}
		parser.next()

		return exp, nil
	}

	// otherwise, try to parse a variable
	variable, err := parser.parseVariable()
	if err != nil {
		return nil, err
	}

	return variable, nil
}
