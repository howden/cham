package ast

// The OR boolean expression - ||
func BooleanOr(left BooleanTerm, right BooleanTerm) BooleanTerm {
	return &booleanOr{left, right}
}

// The AND boolean expression - &&
func BooleanAnd(left BooleanTerm, right BooleanTerm) BooleanTerm {
	return &booleanAnd{left, right}
}

// The NOT boolean expression - !
func BooleanNot(term BooleanTerm) BooleanTerm {
	return &booleanNot{term}
}

type booleanOr struct {
	left  BooleanTerm
	right BooleanTerm
}

type booleanAnd struct {
	left  BooleanTerm
	right BooleanTerm
}

type booleanNot struct {
	exp BooleanTerm
}

func (b booleanOr) Eval(state State) bool {
	return b.left.Eval(state) || b.right.Eval(state)
}

func (b booleanAnd) Eval(state State) bool {
	return b.left.Eval(state) && b.right.Eval(state)
}

func (b booleanNot) Eval(state State) bool {
	return !b.exp.Eval(state)
}
