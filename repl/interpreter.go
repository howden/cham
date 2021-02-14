package repl

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/eval"
	"github.com/howden/cham/lexer"
	"github.com/howden/cham/parser"
)

// Parses a program
func ParseProgram(src string) (*ast.Program, error) {
	return parser.NewParser(lexer.FromString(src)).ParseProgramFully()
}

// Runs a program and prints the result to STDOUT
func PrintEvalOutput(src string) {
	program, err := ParseProgram(src)
	if err != nil {
		parser.PrintParserError(src, err)
		return
	}

	result, err := eval.Evaluate(program)
	if err != nil {
		fmt.Printf("error evaluating: %s\n", err)
		return
	}

	fmt.Println(result)
}

// Runs a program through the parser and prints the resultant AST
func PrintParserOutput(src string) {
	fmt.Println("Parser Output:")
	prog, err := ParseProgram(src)
	if err != nil {
		parser.PrintParserError(src, err)
	} else {
		fmt.Println(prog)
	}
}

// Runs a program through the lexer and prints the resultant tokens
func PrintLexerOutput(src string) {
	fmt.Println("Lexer Output:")
	tokens, err := lexer.FromString(src).RemainingTokens()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, tok := range tokens {
			fmt.Println(tok)
		}
	}
}
