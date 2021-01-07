package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/lexer"
	"github.com/howden/cham/token"
	"strconv"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer}

	// advance twice to fill both current and peek variables
	p.advance()
	p.advance()

	return p
}

func (parser *Parser) ParseProgram() (*ast.Reaction, error) {
	reaction, err := parser.parseReaction()
	if err != nil {
		return nil, err
	}

	_, err = parser.expectToken(token.EOF)
	if err != nil {
		return nil, err
	}

	return reaction, nil
}

func (parser *Parser) advance() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func expect(token token.Token, expected token.TokenType) (bool, error) {
	if token.Type != expected {
		return false, fmt.Errorf("expected token %v but got %v instead", expected, token.Type)
	}
	return true, nil
}

func (parser *Parser) expectToken(expected token.TokenType) (bool, error) {
	return expect(parser.currentToken, expected)
}

func (parser *Parser) parseVariable() (ast.IntegerTerm, error) {
	if ok, _ := parser.expectToken(token.Number); ok {
		i, err := strconv.Atoi(parser.currentToken.Literal)
		if err != nil {
			return nil, fmt.Errorf("parser error for int: %v, %w", parser.currentToken.Literal, err)
		}

		parser.advance()
		return ast.Number(i), nil
	}

	if ok, _ := parser.expectToken(token.Ident); ok {
		parser.advance()
		return ast.Ident(parser.currentToken.Literal), nil
	}

	return nil, fmt.Errorf("expected number or ident but got %v instead", parser.currentToken)
}

func (parser *Parser) parseReaction() (*ast.Reaction, error) {
	// parse input
	input, err := parser.parseReactionInput()
	if err != nil {
		return nil, err
	}

	// expect reaction-op
	if ok, err := parser.expectToken(token.ReactionOp); !ok {
		return nil, err
	}
	parser.advance()

	action, err := parser.parseReactionAction()
	if err != nil {
		return nil, err
	}

	// expect if
	if ok, err := parser.expectToken(token.If); !ok {
		return nil, err
	}
	parser.advance()

	condition, err := parser.parseBexp()
	if err != nil {
		return nil, err
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
			parser.advance()
		}

		// if first=true, or if there was a comma before, require that the next
		// token is an ident, if not return an error
		if ok, err := parser.expectToken(token.Ident); !ok {
			return nil, err
		}

		// append the identifier to the slice and advance the parser
		identifiers = append(identifiers, ast.Ident(parser.currentToken.Literal))
		parser.advance()
	}

	return &ast.ReactionInput{Idents: identifiers}, nil
}

func (parser *Parser) parseReactionAction() (*ast.ReactionAction, error) {
	openCurly, _ := parser.expectToken(token.OpenCurlyBracket)
	if openCurly {
		parser.advance()
	}

	// if immediately closed, return an empty products slice
	if closeCurly, _ := parser.expectToken(token.CloseCurlyBracket); closeCurly {
		parser.advance()
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
			parser.advance()
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
