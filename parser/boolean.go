package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
)

// arithmetic.go contains the parsing code for boolean expressions.
//
// Context-free grammar accepted by this parser:
//   <bexp> ::= <bterm> {<or> <bterm>}
//   <bterm> ::= <bnotfactor> {<and> <bnotfactor>}
//   <bnotfactor> ::= <not> <bfactor>
//   <bnotfactor> ::= <bfactor>
//   <bfactor> ::= <bool-value>
//   <bfactor> ::= <openb> <bexp> <closeb>

// Parses a "bexp"
// <bexp> ::= <bterm> {<or> <bterm>}
func (parser *Parser) parseBexp() (ast.BooleanTerm, error) {
	root, err := parser.bterm()
	if err != nil {
		return nil, err
	}

	for parser.currentToken.Type == token.Or {
		parser.next()

		// use root as the left value, parse right from the next token
		left := root
		right, err := parser.bterm()

		if err != nil {
			return nil, err
		}

		// set root to the result
		root = ast.BooleanOr(left, right)
	}

	return root, nil
}

// Parses a "bterm"
// <bterm> ::= <bnotfactor> {<and> <bnotfactor>}
func (parser *Parser) bterm() (ast.BooleanTerm, error) {
	root, err := parser.notFactor()
	if err != nil {
		return nil, err
	}

	for parser.currentToken.Type == token.And {
		parser.next()
		// use root as the left value, parse right from the next token
		left := root
		right, err := parser.notFactor()

		if err != nil {
			return nil, err
		}

		// set root to the result
		root = ast.BooleanAnd(left, right)
	}

	return root, nil
}

// Parses a "bnotfactor"
// <bnotfactor> ::= <not> <bfactor>
// <bnotfactor> ::= <bfactor>
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

// Parses a "bfactor"
// <bfactor> ::= <bool-value>
// <bfactor> ::= <openb> <bexp> <closeb>
func (parser *Parser) bfactor() (ast.BooleanTerm, error) {
	// try to parse brackets first
	if parser.currentToken.Type == token.OpenBracket {
		parser.next()

		// parse inner bexp
		exp, err := parser.parseBexp()
		if err != nil {
			return nil, err
		}

		// ensure bracket is closed after bexp is finished
		if parser.currentToken.Type != token.CloseBracket {
			return nil, fmt.Errorf("expected close bracket but got %v instead", parser.currentToken)
		}
		parser.next()

		return exp, nil
	}

	// otherwise, try to parse a comparison
	comp, err := parser.parseComparison()
	if err != nil {
		return nil, err
	}

	return comp, nil
}
