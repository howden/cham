package parser

import (
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
	"github.com/pkg/errors"
)

func (parser *Parser) parseIdentTuple() (*ast.IdentifierTuple, error) {
	openTuple, _ := parser.expectToken(token.OpenSquareBracket)
	if openTuple {
		parser.next()
	}

	var idents []ast.Identifier

	first, err := parser.parseIdent()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing identifier")
	}
	idents = append(idents, ast.Ident(first))

	for openTuple && parser.currentToken.Type == token.Comma {
		parser.next()

		ident, err := parser.parseIdent()
		if err != nil {
			return nil, errors.Wrap(err, "error parsing identifier")
		}
		idents = append(idents, ast.Ident(ident))
	}

	if openTuple {
		ok, err := parser.expectToken(token.CloseSquareBracket)
		if !ok {
			return nil, err
		}
		parser.next()
	}

	return &ast.IdentifierTuple{Values: idents}, nil
}

func (parser *Parser) parseNumberTuple() (*ast.IntTuple, error) {
	openTuple, _ := parser.expectToken(token.OpenSquareBracket)
	if openTuple {
		parser.next()
	}

	var ints []int

	first, err := parser.parseNumber()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing number")
	}
	ints = append(ints, first)

	for openTuple && parser.currentToken.Type == token.Comma {
		parser.next()

		val, err := parser.parseNumber()
		if err != nil {
			return nil, errors.Wrap(err, "error parsing number")
		}
		ints = append(ints, val)
	}

	if openTuple {
		ok, err := parser.expectToken(token.CloseSquareBracket)
		if !ok {
			return nil, err
		}
		parser.next()
	}

	tuple := ast.CreateIntTuple(ints)
	return &tuple, nil
}

func (parser *Parser) parseAexpTuple() (*ast.IntegerTermTuple, error) {
	openTuple, _ := parser.expectToken(token.OpenSquareBracket)
	if openTuple {
		parser.next()
	}

	var vars []ast.IntegerTerm

	first, err := parser.parseAexp()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing arithmetic expression")
	}
	vars = append(vars, first)

	for openTuple && parser.currentToken.Type == token.Comma {
		parser.next()

		variable, err := parser.parseAexp()
		if err != nil {
			return nil, errors.Wrap(err, "error parsing arithmetic expression")
		}
		vars = append(vars, variable)
	}

	if openTuple {
		ok, err := parser.expectToken(token.CloseSquareBracket)
		if !ok {
			return nil, err
		}
		parser.next()
	}

	return &ast.IntegerTermTuple{Values: vars}, nil
}
