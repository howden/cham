package token

import "fmt"

type TokenType int

const (
	EOF TokenType = iota
	Invalid
	Ident
	Number
	ReactionChain      // |
	ReactionDef        // :
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
	Modulo             // %
	Comma              // ,
	OpenBracket        // (
	CloseBracket       // )
	OpenCurlyBracket   // {
	CloseCurlyBracket  // }
	OpenSquareBracket  // [
	CloseSquareBracket // ]
)

func (t TokenType) New() Token {
	return Token{Type: t}
}

func (t TokenType) WithLiteral(literal string) Token {
	return Token{Type: t, Literal: literal, hasLiteral: true}
}

func Error(err error) Token {
	return Token{Type: Invalid, Err: err}
}

var names = map[TokenType]string{
	EOF:                "EOF",
	Invalid:            "Invalid",
	Ident:              "ident",
	Number:             "number",
	ReactionChain:      "reactionChain",
	ReactionDef:        "reactionDef",
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
	Modulo:             "modulo",
	Comma:              "comma",
	OpenBracket:        "openBracket",
	CloseBracket:       "closeBracket",
	OpenCurlyBracket:   "openCurlyBracket",
	CloseCurlyBracket:  "closeCurlyBracket",
	OpenSquareBracket:  "openSquareBracket",
	CloseSquareBracket: "closeSquareBracket",
}

func (t TokenType) String() string {
	return names[t]
}

type Token struct {
	Type       TokenType
	Literal    string
	hasLiteral bool
	Err        error
}

func (t Token) String() string {
	if t.Type == Invalid && t.Err != nil {
		return fmt.Sprintf("invalid(%s)", t.Err)
	}
	if t.hasLiteral {
		return fmt.Sprintf("%s(%s)", t.Type, t.Literal)
	}
	return t.Type.String()
}
