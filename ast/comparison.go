package ast

import (
	"fmt"
)

// Type representing a comparison of two integers.
type Comparison struct {
	left         IntegerTerm
	right        IntegerTerm
	operator     func(left int, right int) bool
	operatorName string
}

func (c Comparison) Eval(state State) (bool, error) {
	l, err := c.left.Eval(state)
	if err != nil {
		return false, err
	}

	r, err := c.right.Eval(state)
	if err != nil {
		return false, err
	}

	return c.operator(l, r), nil
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

func equals(left int, right int) bool {
	return left == right
}

func notEquals(left int, right int) bool {
	return left != right
}

func lessThan(left int, right int) bool {
	return left < right
}

func greaterThan(left int, right int) bool {
	return left > right
}

func lessThanEqual(left int, right int) bool {
	return left <= right
}

func greaterThanEqual(left int, right int) bool {
	return left >= right
}
