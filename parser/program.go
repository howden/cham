package parser

import (
	"github.com/howden/cham/ast"
	"github.com/howden/cham/eval"
	"github.com/howden/cham/token"
	"github.com/pkg/errors"
)

// Parses a full program, terminated by EOF
func (parser *Parser) ParseProgramFully() (*ast.Program, error) {
	program, err := parser.parseProgram(nil)
	if err != nil {
		return nil, parser.wrapError(err)
	}

	_, err = parser.expectToken(token.EOF)
	if err != nil {
		return nil, parser.wrapError(err)
	}

	return program, nil
}

// Parses a full program or reaction definition, terminated by EOF
func (parser *Parser) ParseProgramOrReactionDefFully(store *eval.ReactionStore) (*ast.Program, *ast.ReactionPointer, error) {
	var program *ast.Program
	var reaction *ast.ReactionPointer
	var err error

	if parser.currentToken.Type == token.Ident {
		reaction, err = parser.parseReactionDefinition(store)
		if err != nil {
			return nil, nil, parser.wrapError(err)
		}
	} else {
		program, err = parser.parseProgram(store)
		if err != nil {
			return nil, nil, parser.wrapError(err)
		}
	}

	_, err = parser.expectToken(token.EOF)
	if err != nil {
		return nil, nil, parser.wrapError(err)
	}

	return program, reaction, nil
}

func (parser *Parser) parseReactionDefinition(store *eval.ReactionStore) (*ast.ReactionPointer, error) {
	ident, err := parser.parseIdent()
	if err != nil {
		return nil, err
	}

	// expect reaction-def
	if ok, err := parser.expectToken(token.ReactionDef); !ok {
		return nil, err
	}
	parser.next()

	reactions, err := parser.parseReactions(store)
	if err != nil {
		return nil, err
	}

	return &ast.ReactionPointer{
		Identifier: ast.Ident(ident),
		Reactions:  reactions,
	}, nil
}

func (parser *Parser) parseProgram(store *eval.ReactionStore) (*ast.Program, error) {
	input, err := parser.parseInput()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing program input")
	}

	// expect reaction-chain
	if ok, err := parser.expectToken(token.ReactionChain); !ok {
		return nil, err
	}
	parser.next()

	reactions, err := parser.parseReactions(store)
	if err != nil {
		return nil, err
	}

	return &ast.Program{Input: input, Reactions: reactions}, nil
}

func (parser *Parser) parseReactions(store *eval.ReactionStore) ([]*ast.Reaction, error) {
	var reactions []*ast.Reaction

	first, err := parser.parseReactionPointer(store)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing program reaction")
	}
	for _, r := range first {
		reactions = append(reactions, r)
	}

	for parser.currentToken.Type == token.ReactionChain {
		parser.next()

		reaction, err := parser.parseReactionPointer(store)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing program reaction")
		}
		for _, r := range reaction {
			reactions = append(reactions, r)
		}
	}

	return reactions, nil
}

func (parser *Parser) parseInput() ([]ast.IntTuple, error) {
	openCurly, _ := parser.expectToken(token.OpenCurlyBracket)
	if openCurly {
		parser.next()

		// if immediately closed, return an empty input slice
		if closeCurly, _ := parser.expectToken(token.CloseCurlyBracket); closeCurly {
			parser.next()
			return []ast.IntTuple{}, nil
		}
	}

	// otherwise, expect numbers separated by commas
	var ints []ast.IntTuple

	first, err := parser.parseNumberTuple()
	if err != nil {
		return nil, err
	}
	ints = append(ints, *first)

	// keep accepting more numbers while there are commas
	for parser.currentToken.Type == token.Comma {
		parser.next()

		num, err := parser.parseNumberTuple()
		if err != nil {
			return nil, err
		}
		ints = append(ints, *num)
	}

	if openCurly {
		ok, err := parser.expectToken(token.CloseCurlyBracket)
		if !ok {
			return nil, err
		}
		parser.next()
	}

	return ints, nil
}
