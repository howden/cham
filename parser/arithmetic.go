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

// Map of "addop" tokens -> a function that creates an AST
var addOps = map[token.TokenType]func(left ast.IntegerTerm, right ast.IntegerTerm) ast.ArithmeticExp{
	token.Plus:     ast.Plus,
	token.Subtract: ast.Subtract,
}

// Map of "multop" tokens -> a function that creates an AST
var multOps = map[token.TokenType]func(left ast.IntegerTerm, right ast.IntegerTerm) ast.ArithmeticExp{
	token.Multiply: ast.Multiply,
	token.Divide:   ast.Divide,
}

// Parses an "aexp"
// <aexp> ::= <aterm> {<addop> <aexp>}
func (parser *Parser) parseAexp() (ast.IntegerTerm, error) {
	root, err := parser.aterm()
	if err != nil {
		return nil, err
	}

	for {
		operation, ok := addOps[parser.currentToken.Type]
		if !ok {
			break
		}
		parser.next()

		// use root as the left value, parse right from the next token
		left := root
		right, err := parser.parseAexp()

		if err != nil {
			return nil, err
		}

		root = operation(left, right)
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

	for {
		operation, ok := multOps[parser.currentToken.Type]
		if !ok {
			break
		}
		parser.next()

		// use root as the left value, parse right from the next token
		left := root
		right, err := parser.aterm()

		if err != nil {
			return nil, err
		}

		root = operation(left, right)
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
