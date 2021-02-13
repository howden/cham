package ast

import (
	"fmt"
)

// Type representing an arithmetic expression between two integer terms.
// Since this produces an int result, an ArithmeticExp is also an integer term.
type ArithmeticExp struct {
	left         IntegerTerm
	right        IntegerTerm
	operator     func(left int, right int) int
	operatorName string
}

func (a ArithmeticExp) Eval(state State) (int, error) {
	l, err := a.left.Eval(state)
	if err != nil {
		return 0, err
	}

	r, err := a.right.Eval(state)
	if err != nil {
		return 0, err
	}

	return a.operator(l, r), nil
}

func (a ArithmeticExp) String() string {
	return fmt.Sprintf("%s{%v, %v}", a.operatorName, a.left, a.right)
}

// Returns a 'plus' arithmetic expression between the two given terms
func Plus(left IntegerTerm, right IntegerTerm) ArithmeticExp {
	return ArithmeticExp{left, right, plus, "plus"}
}

// Returns a 'subtract' arithmetic expression between the two given terms
func Subtract(left IntegerTerm, right IntegerTerm) ArithmeticExp {
	return ArithmeticExp{left, right, subtract, "subtract"}
}

// Returns a 'multiply' arithmetic expression between the two given terms
func Multiply(left IntegerTerm, right IntegerTerm) ArithmeticExp {
	return ArithmeticExp{left, right, multiply, "multiply"}
}

// Returns a 'divide' arithmetic expression between the two given terms
func Divide(left IntegerTerm, right IntegerTerm) ArithmeticExp {
	return ArithmeticExp{left, right, divide, "divide"}
}

// Returns a 'modulo' arithmetic expression between the two given terms
func Modulo(left IntegerTerm, right IntegerTerm) ArithmeticExp {
	return ArithmeticExp{left, right, modulo, "modulo"}
}

func plus(left int, right int) int {
	return left + right
}

func subtract(left int, right int) int {
	return left - right
}

func multiply(left int, right int) int {
	return left * right
}

func divide(left int, right int) int {
	return left / right
}

func modulo(left int, right int) int {
	return left % right
}
