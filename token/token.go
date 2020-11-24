package token

import "fmt"

type TokenType int

const (
	EOF TokenType = iota
	Ident
	Number
	ReactionOp         // =>
	If                 // if
	LessThan           // <
	GreaterThan        // >
	LessThanOrEqual    // <=
	GreaterThanOrEqual // >=
	Equal              // ==
	NotEqual           // !=
	Not                // !
	Or                 // ||
	And                // &&
	Plus               // +
	Subtract           // -
	Multiply           // *
	Divide             // /
	Comma              // ,
	OpenBracket        // (
	CloseBracket       // )
	OpenCurlyBracket   // {
	CloseCurlyBracket  // }
)

func (t TokenType) New() Token {
	return Token{Type: t}
}

func (t TokenType) WithLiteral(literal string) Token {
	return Token{Type: t, Literal: &literal}
}

var names = map[TokenType]string{
	EOF:                "EOF",
	Ident:              "ident",
	Number:             "number",
	ReactionOp:         "reactionOp",
	If:                 "if",
	LessThan:           "lessThan",
	GreaterThan:        "greaterThan",
	LessThanOrEqual:    "lessThanOrEqual",
	GreaterThanOrEqual: "greaterThanOrEqual",
	Equal:              "equal",
	NotEqual:           "notEqual",
	Not:                "not",
	Or:                 "or",
	And:                "and",
	Plus:               "plus",
	Subtract:           "subtract",
	Multiply:           "multiply",
	Divide:             "divide",
	Comma:              "comma",
	OpenBracket:        "openBracket",
	CloseBracket:       "closeBracket",
	OpenCurlyBracket:   "openCurlyBracket",
	CloseCurlyBracket:  "closeCurlyBracket",
}

func (t TokenType) String() string {
	return names[t]
}

type Token struct {
	Type    TokenType
	Literal *string
}

func (t Token) String() string {
	if t.Literal != nil {
		return fmt.Sprintf("%s(%s)", t.Type, *t.Literal)
	}
	return fmt.Sprintf("%s", t.Type)
}
