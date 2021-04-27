package repl

import (
	"fmt"
	"github.com/howden/cham/ast"
	"github.com/howden/cham/eval"
	"github.com/howden/cham/lexer"
	"github.com/howden/cham/parser"
	"strings"
)

// Parses a program
func ParseProgram(src string) (*ast.Program, error) {
	return parser.NewParser(lexer.FromString(src)).ParseProgramFully()
}

// Parses a program or reaction definition
func ParseProgramOrReaction(src string, store *eval.ReactionStore) (*ast.Program, *ast.ReactionPointer, error) {
	return parser.NewParser(lexer.FromString(src)).ParseProgramOrReactionDefFully(store)
}

// Runs a program from the REPL and prints the result to STDOUT
func HandleReplInput(src string, store *eval.ReactionStore) {
	program, reactionDef, err := ParseProgramOrReaction(src, store)
	if err != nil {
		parser.PrintParserError(src, err)
		return
	}

	if program != nil {
		result, err := eval.Evaluate(program)
		if err != nil {
			fmt.Printf("error evaluating: %s\n", err)
		} else {
			fmt.Println(result)
		}
	} else if reactionDef != nil {
		store.Put(reactionDef)
		fmt.Println("OK")
	}
}

// Runs a program and prints the result to STDOUT
func HandleCmdLineInput(src string) {
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

// Runs a program loaded from a file and prints the result to STDOUT
func HandleFileInput(lines []string, store *eval.ReactionStore) {
	for _, src := range lines {
		src = strings.Trim(src, " \n\t")
		if len(src) == 0 {
			continue
		}

		program, reactionDef, err := ParseProgramOrReaction(src, store)
		if err != nil {
			parser.PrintParserError(src, err)
			return
		}

		if program != nil {
			result, err := eval.Evaluate(program)
			if err != nil {
				fmt.Printf("error evaluating: %s\n", err)
				return
			} else {
				fmt.Println(result)
			}
		} else if reactionDef != nil {
			store.Put(reactionDef)
		}
	}
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
