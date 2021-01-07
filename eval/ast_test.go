package eval

import (
	"github.com/howden/cham/ast"
	"testing"
)

func TestAexp(t *testing.T) {
	state := NewState()
	state.PutVar(ast.Ident("x"), 1)

	// x + 2  ==>  3
	aexp := ast.Plus(ast.Ident("x"), ast.Number(2))

	result := aexp.Eval(state)
	if result != 3 {
		t.Errorf("expected 3 but got %d", result)
	}
}

func TestBexp(t *testing.T) {
	state := NewState()
	state.PutVar(ast.Ident("x"), 1)

	// (x == 1) && !(12 != 12)  ==>  true
	bexp := ast.BooleanAnd(
		ast.Equals(ast.Ident("x"), ast.Number(1)),
		ast.BooleanNot(ast.NotEquals(ast.Number(12), ast.Number(12))))

	result := bexp.Eval(state)
	if !result {
		t.Errorf("expected true but got false")
	}
}

func TestBexp2(t *testing.T) {
	state := NewState()
	state.PutVar(ast.Ident("x"), 1)

	// (x - 2) > -5  ==>  true
	bexp := ast.GreaterThan(ast.Subtract(ast.Ident("x"), ast.Number(2)), ast.Number(-5))

	result := bexp.Eval(state)
	if !result {
		t.Errorf("expected true but got false")
	}
}
