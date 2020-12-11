package main

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/eval"
	"github.com/howden/cham/lexer"
	"os"
	"strings"
)

func textLexer() {
	if len(os.Args) < 2 {
		fmt.Println("./cham <program>")
		return
	}

	src := strings.Join(os.Args[1:], " ")
	tokens, err := lexer.FromString(src).RemainingTokens()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, tok := range tokens {
		fmt.Println(tok)
	}
}

func main() {
	// test ast & simple evaluation

	state := eval.NewState()
	state.PutVar(ast.Ident("x"), 1)

	// x + 2  ==>  3
	aexp := ast.Plus(ast.Ident("x"), ast.Number(2))
	fmt.Println(aexp.Eval(state))

	// (x == 1) && !(12 != 12)  ==>  true
	bexp := ast.BooleanAnd(
		ast.Equals(ast.Ident("x"), ast.Number(1)),
		ast.BooleanNot(ast.NotEquals(ast.Number(12), ast.Number(12))))
	fmt.Println(bexp.Eval(state))

	// (x - 2) > -5  ==>  true
	bexp2 := ast.GreaterThan(ast.Subtract(ast.Ident("x"), ast.Number(2)), ast.Number(-5))
	fmt.Println(bexp2.Eval(state))
}
