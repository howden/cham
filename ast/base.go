package ast

// Interface to represent the program state.
// (stores the values of variables in a given reaction)
type State interface {
	GetVar(ident *Identifier) int
}

// Interface representing an integer term - just something that returns an int
// Variables could be:
// - an Identifier
// - a number
// - an ArithmeticExp
type IntegerTerm interface {
	Eval(state State) int
}

// Type representing a boolean term.
// Could be:
// - a Comparison
// - a boolean expression - BooleanOr, BooleanAnd, BooleanNot
type BooleanTerm interface {
	Eval(state State) bool
}

// Creates an identifier
func Ident(s string) *Identifier {
	return &Identifier{s}
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

func (ident *Identifier) Eval(state State) int {
	return state.GetVar(ident)
}

func (ident *Identifier) String() string {
	return ident.name
}

// Struct representing a number
// These are used in the action/condition
type number struct {
	Int int
}

func (number number) Eval(_ State) int {
	return number.Int
}
