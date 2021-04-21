package ast

import (
	"fmt"
)

// Interface to represent the program state.
// (stores the values of variables in a given reaction)
type State interface {
	GetVar(ident Identifier) (int, error)
}

// Interface representing an integer term - just something that returns an int
// Variables could be:
// - an Identifier
// - a number
// - an ArithmeticExp
type IntegerTerm interface {
	Eval(state State) (int, error)
}

// Type representing a boolean term.
// Could be:
// - a Comparison
// - a boolean expression - BooleanOr, BooleanAnd, BooleanNot
type BooleanTerm interface {
	Eval(state State) (bool, error)
}

// Creates an identifier
func Ident(s string) Identifier {
	return Identifier{name: s}
}

// Creates a number
func Number(i int) IntegerTerm {
	return &number{i}
}

// Struct representing an identifier
// These are defined by the input, then referenced in the action/condition.
type Identifier struct {
	name string
}

func (ident Identifier) Eval(state State) (int, error) {
	return state.GetVar(ident)
}

func (ident Identifier) String() string {
	return fmt.Sprintf("ident(%s)", ident.name)
}

func (ident Identifier) Name() string {
	return ident.name
}

// Struct representing a number
// These are used in the action/condition
type number struct {
	int int
}

func (number number) Eval(_ State) (int, error) {
	return number.int, nil
}

func (number number) String() string {
	return fmt.Sprintf("number(%v)", number.int)
}
