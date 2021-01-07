package ast

import (
	"fmt"
)

// Type representing a comparison of two integers.
type Comparison struct {
	left         IntegerTerm
	right        IntegerTerm
	operator     func(state State, left IntegerTerm, right IntegerTerm) bool
	operatorName string
}

func (c Comparison) Eval(state State) bool {
	return c.operator(state, c.left, c.right)
}

func (c Comparison) String() string {
	return fmt.Sprintf("%s{%v, %v}", c.operatorName, c.left, c.right)
}

// Returns an 'equal to' comparison between the two given terms
func Equals(left IntegerTerm, right IntegerTerm) BooleanTerm {
	return &Comparison{left, right, equals, "equals"}
}

// Returns a 'not equal to' comparison between the two given terms
func NotEquals(left IntegerTerm, right IntegerTerm) BooleanTerm {
	return &Comparison{left, right, notEquals, "notEquals"}
}

// Returns a 'less than' comparison between the two given terms
func LessThan(left IntegerTerm, right IntegerTerm) BooleanTerm {
	return &Comparison{left, right, lessThan, "lessThan"}
}

// Returns a 'greater than' comparison between the two given terms
func GreaterThan(left IntegerTerm, right IntegerTerm) BooleanTerm {
	return &Comparison{left, right, greaterThan, "greaterThan"}
}

// Returns a 'less than or equal to' comparison between the two given terms
func LessThanEqual(left IntegerTerm, right IntegerTerm) BooleanTerm {
	return &Comparison{left, right, lessThanEqual, "lessThanEqual"}
}

// Returns a 'greater than or equal to' comparison between the two given terms
func GreaterThanEqual(left IntegerTerm, right IntegerTerm) BooleanTerm {
	return &Comparison{left, right, greaterThanEqual, "greaterThanEqual"}
}

func equals(state State, left IntegerTerm, right IntegerTerm) bool {
	return left.Eval(state) == right.Eval(state)
}

func notEquals(state State, left IntegerTerm, right IntegerTerm) bool {
	return left.Eval(state) != right.Eval(state)
}

func lessThan(state State, left IntegerTerm, right IntegerTerm) bool {
	return left.Eval(state) < right.Eval(state)
}

func greaterThan(state State, left IntegerTerm, right IntegerTerm) bool {
	return left.Eval(state) > right.Eval(state)
}

func lessThanEqual(state State, left IntegerTerm, right IntegerTerm) bool {
	return left.Eval(state) <= right.Eval(state)
}

func greaterThanEqual(state State, left IntegerTerm, right IntegerTerm) bool {
	return left.Eval(state) >= right.Eval(state)
}
