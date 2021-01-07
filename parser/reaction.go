package parser

import (
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
	"github.com/pkg/errors"
)

// Parses a full "reaction" program, terminated by EOF
func (parser *Parser) ParseProgram() (*ast.Reaction, error) {
	reaction, err := parser.parseReaction()
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing, at position: %s", parser.lexer.Pos())
	}

	_, err = parser.expectToken(token.EOF)
	if err != nil {
		return nil, err
	}

	return reaction, nil
}

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

	// expect if
	if ok, err := parser.expectToken(token.If); !ok {
		return nil, err
	}
	parser.next()

	condition, err := parser.parseBexp()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing reaction condition")
	}

	return &ast.Reaction{
		Input:     input,
		Action:    action,
		Condition: &ast.ReactionCondition{Expression: condition},
	}, nil
}

func (parser *Parser) parseReactionInput() (*ast.ReactionInput, error) {
	var identifiers []ast.Identifier
	first := true

	// keep accepting more identifiers while there are commas
	for next, _ := parser.expectToken(token.Comma); first || next; {
		if first {
			first = false
		} else {
			parser.next()
		}

		// if first=true, or if there was a comma before, require that the next
		// token is an ident, if not return an error
		if ok, err := parser.expectToken(token.Ident); !ok {
			return nil, err
		}

		// append the identifier to the slice and advance the parser
		identifiers = append(identifiers, ast.Ident(parser.currentToken.Literal))
		parser.next()
	}

	return &ast.ReactionInput{Idents: identifiers}, nil
}

func (parser *Parser) parseReactionAction() (*ast.ReactionAction, error) {
	openCurly, _ := parser.expectToken(token.OpenCurlyBracket)
	if openCurly {
		parser.next()
	}

	// if immediately closed, return an empty products slice
	if closeCurly, _ := parser.expectToken(token.CloseCurlyBracket); closeCurly {
		parser.next()
		return &ast.ReactionAction{Products: []ast.IntegerTerm{}}, nil
	}

	// otherwise, expect aexps separated by commas
	var aexps []ast.IntegerTerm
	first := true

	// keep accepting more aexps while there are commas
	for next, _ := parser.expectToken(token.Comma); first || next; {
		if first {
			first = false
		} else {
			parser.next()
		}

		// if first=true, or if there was a comma before, require that the next
		// token is an ident, if not return an error
		aexp, err := parser.parseAexp()
		if err != nil {
			return nil, err
		}
		aexps = append(aexps, aexp)
	}

	return &ast.ReactionAction{Products: aexps}, nil
}
