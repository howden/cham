package parser

import (
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
	"github.com/pkg/errors"
)

// Parses a full program, terminated by EOF
func (parser *Parser) ParseProgramFully() (*ast.Program, error) {
	program, err := parser.parseProgram()
	if err != nil {
		return nil, parser.wrapError(err)
	}

	_, err = parser.expectToken(token.EOF)
	if err != nil {
		return nil, parser.wrapError(err)
	}

	return program, nil
}

func (parser *Parser) parseProgram() (*ast.Program, error) {
	input, err := parser.parseInput()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing program input")
	}

	// expect reaction-chain
	if ok, err := parser.expectToken(token.ReactionChain); !ok {
		return nil, err
	}
	parser.next()

	var reactions []*ast.Reaction

	first, err := parser.parseReaction()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing program reaction")
	}
	reactions = append(reactions, first)

	for parser.currentToken.Type == token.ReactionChain {
		parser.next()

		reaction, err := parser.parseReaction()
		if err != nil {
			return nil, errors.Wrap(err, "error parsing program reaction")
		}
		reactions = append(reactions, reaction)
	}

	return &ast.Program{Input: input, Reactions: reactions}, nil
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
