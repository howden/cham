package main

import (
	"fmt"
	"github.com/howden/cham/eval"
	"github.com/howden/cham/lexer"
	"github.com/howden/cham/parser"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("./cham <program>")
		return
	}

	src := strings.Join(os.Args[1:], " ")

	program, err := parser.NewParser(lexer.FromString(src)).ParseProgramFully()
	if err != nil {
		parser.PrintParserError(src, err)
		return
	}

	result := eval.Evaluate(program)
	fmt.Println(result)
}

/*
func demoLexer(src string) {
	fmt.Println("Lexer Output:")
	tokens, err := lexer.FromString(src).RemainingTokens()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, tok := range tokens {
			fmt.Println(tok)
		}
	}
	fmt.Print("\n\n")
}

func demoParser(src string) {
	fmt.Println("Parser Output:")
	ast, err := parser.NewParser(lexer.FromString(src)).ParseProgramFully()
	if err != nil {
		parser.PrintParserError(src, err)
	} else {
		fmt.Println(ast)
	}
}
*/
