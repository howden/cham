package ast

// Type representing an arithmetic expression between two integer terms.
// Since this produces an int result, an ArithmeticExp is also an integer term.
type ArithmeticExp struct {
	left     IntegerTerm
	right    IntegerTerm
	operator func(state State, left IntegerTerm, right IntegerTerm) int
}

func (a *ArithmeticExp) Eval(state State) int {
	return a.operator(state, a.left, a.right)
}

// Returns a 'plus' arithmetic expression between the two given terms
func Plus(left IntegerTerm, right IntegerTerm) *ArithmeticExp {
	return &ArithmeticExp{left, right, plus}
}

// Returns a 'subtract' arithmetic expression between the two given terms
func Subtract(left IntegerTerm, right IntegerTerm) *ArithmeticExp {
	return &ArithmeticExp{left, right, subtract}
}

// Returns a 'multiply' arithmetic expression between the two given terms
func Multiply(left IntegerTerm, right IntegerTerm) *ArithmeticExp {
	return &ArithmeticExp{left, right, multiply}
}

// Returns a 'divide' arithmetic expression between the two given terms
func Divide(left IntegerTerm, right IntegerTerm) *ArithmeticExp {
	return &ArithmeticExp{left, right, divide}
}

func plus(state State, left IntegerTerm, right IntegerTerm) int {
	return left.Eval(state) + right.Eval(state)
}

func subtract(state State, left IntegerTerm, right IntegerTerm) int {
	return left.Eval(state) - right.Eval(state)
}

func multiply(state State, left IntegerTerm, right IntegerTerm) int {
	return left.Eval(state) * right.Eval(state)
}

func divide(state State, left IntegerTerm, right IntegerTerm) int {
	return left.Eval(state) / right.Eval(state)
}
