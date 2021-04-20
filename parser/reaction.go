package parser

import (
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
	"github.com/pkg/errors"
)

func (parser *Parser) parseReaction() (*ast.Reaction, error) {
	// parse input
	input, err := parser.parseReactionInput()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing reaction input")
	}

	// expect reaction-op
	if ok, err := parser.expectToken(token.ReactionOp); !ok {
		return nil, err
	}
	parser.next()

	action, err := parser.parseReactionAction()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing reaction action")
	}

	condition, err := parser.parseReactionCondition()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing reaction condition")
	}

	return &ast.Reaction{
		Input:     input,
		Action:    action,
		Condition: condition,
	}, nil
}

func (parser *Parser) parseReactionInput() (*ast.ReactionInput, error) {
	var identifiers []ast.IdentifierTuple

	first, err := parser.parseIdentTuple()
	if err != nil {
		return nil, err
	}
	identifiers = append(identifiers, *first)

	// keep accepting more identifiers while there are commas
	for parser.currentToken.Type == token.Comma {
		parser.next()

		ident, err := parser.parseIdentTuple()
		if err != nil {
			return nil, err
		}

		// append the identifier to the slice and advance the parser
		identifiers = append(identifiers, *ident)
	}

	return &ast.ReactionInput{Idents: identifiers}, nil
}

func (parser *Parser) parseReactionAction() (*ast.ReactionAction, error) {
	openCurly, _ := parser.expectToken(token.OpenCurlyBracket)
	if openCurly {
		parser.next()

		// if immediately closed, return an empty products slice
		if closeCurly, _ := parser.expectToken(token.CloseCurlyBracket); closeCurly {
			parser.next()
			return &ast.ReactionAction{Products: []ast.IntegerTermTuple{}}, nil
		}
	}

	// otherwise, expect aexps separated by commas
	var aexps []ast.IntegerTermTuple

	first, err := parser.parseAexpTuple()
	if err != nil {
		return nil, err
	}
	aexps = append(aexps, *first)

	// keep accepting more aexps while there are commas
	for parser.currentToken.Type == token.Comma {
		parser.next()

		aexp, err := parser.parseAexpTuple()
		if err != nil {
			return nil, err
		}
		aexps = append(aexps, *aexp)
	}

	if openCurly {
		ok, err := parser.expectToken(token.CloseCurlyBracket)
		if !ok {
			return nil, err
		}
		parser.next()
	}

	return &ast.ReactionAction{Products: aexps}, nil
}

func (parser *Parser) parseReactionCondition() (*ast.ReactionCondition, error) {
	if ok, _ := parser.expectToken(token.If); !ok {
		return &ast.ReactionCondition{Expression: ast.BooleanTrue()}, nil
	}
	parser.next()

	condition, err := parser.parseBexp()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing reaction condition")
	}

	return &ast.ReactionCondition{Expression: condition}, nil
}
