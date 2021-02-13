package ast

import (
	"fmt"
)

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

// Boolean expression that always evaluates to true
func BooleanTrue() BooleanTerm {
	return &booleanConst{true}
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

type booleanConst struct {
	val bool
}

func (b booleanOr) Eval(state State) (bool, error) {
	l, err := b.left.Eval(state)
	if err != nil {
		return false, err
	}

	r, err := b.right.Eval(state)
	if err != nil {
		return false, err
	}

	return l || r, nil
}

func (b booleanOr) String() string {
	return fmt.Sprintf("boolOr{%v, %v}", b.left, b.right)
}

func (b booleanAnd) Eval(state State) (bool, error) {
	l, err := b.left.Eval(state)
	if err != nil {
		return false, err
	}

	r, err := b.right.Eval(state)
	if err != nil {
		return false, err
	}

	return l && r, nil
}

func (b booleanAnd) String() string {
	return fmt.Sprintf("boolAnd{%v, %v}", b.left, b.right)
}

func (b booleanNot) Eval(state State) (bool, error) {
	exp, err := b.exp.Eval(state)
	if err != nil {
		return false, err
	}

	return !exp, nil
}

func (b booleanNot) String() string {
	return fmt.Sprintf("boolNot{%v}", b.exp)
}

func (b booleanConst) Eval(_ State) (bool, error) {
	return b.val, nil
}

func (b booleanConst) String() string {
	if b.val {
		return "boolTrue"
	} else {
		return "boolFalse"
	}
}
