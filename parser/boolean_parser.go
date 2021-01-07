package parser

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/token"
)

type booleanExpressionParser struct {
	parser       *Parser
	currentToken token.Token
	root         ast.BooleanTerm
}

func (parser *Parser) parseBexp() (ast.BooleanTerm, error) {
	boolParser := booleanExpressionParser{parser: parser}
	return boolParser.parse()
}

// <bexp> ::= <bterm> {<or> <bterm>}
// <bterm> ::= <bnotfactor> {<and> <bnotfactor>}
// <bnotfactor> ::= <not> <bfactor>
// <bnotfactor> ::= <bfactor>
// <bfactor> ::= <bool-value>
// <bfactor> ::= <openb> <bexp> <closeb>

func (parser *booleanExpressionParser) parse() (ast.BooleanTerm, error) {
	err := parser.exp()
	if err != nil {
		return nil, err
	}
	return parser.root, nil
}

func (parser *booleanExpressionParser) exp() error {
	err := parser.term()
	if err != nil {
		return err
	}

	for parser.currentToken.Type == token.Or {
		left := parser.root
		err = parser.term()
		if err != nil {
			return err
		}
		right := parser.root
		parser.root = ast.BooleanOr(left, right)
	}

	return nil
}

func (parser *booleanExpressionParser) term() error {
	err := parser.factor()
	if err != nil {
		return err
	}

	for parser.currentToken.Type == token.And {
		left := parser.root
		err = parser.factor()
		if err != nil {
			return err
		}
		right := parser.root
		parser.root = ast.BooleanAnd(left, right)
	}

	return nil
}

func (parser *booleanExpressionParser) notFactor() error {
	parser.currentToken = parser.parser.currentToken

	if parser.currentToken.Type != token.Not {
		return parser.factor()
	}

	parser.parser.advance()

	err := parser.factor()
	if err != nil {
		return err
	}

	parser.root = ast.BooleanNot(parser.root)
	return nil
}

func (parser *booleanExpressionParser) factor() error {
	parser.currentToken = parser.parser.currentToken

	if parser.currentToken.Type == token.OpenBracket {
		parser.parser.advance()
		err := parser.exp()
		if err != nil {
			return err
		}
		if parser.currentToken.Type != token.CloseBracket {
			return fmt.Errorf("expected close bracket but got %v instead", parser.currentToken)
		}
	} else {
		comp, err := parser.parser.parseComparison()
		if err != nil {
			return err
		}
		parser.root = comp
		parser.currentToken = parser.parser.currentToken
		parser.parser.advance()
	}

	return nil
}
