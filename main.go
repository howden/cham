package main

import (
	"fmt"
	"github.com/howden/cham/eval"
	"github.com/howden/cham/lexer"
	"github.com/howden/cham/parser"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		runRepl()
	} else {
		run(strings.Join(os.Args[1:], " "))
	}
}

// Runs a program and outputs the result to STDOUT
func run(src string) {
	program, err := parser.NewParser(lexer.FromString(src)).ParseProgramFully()
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

// Runs the REPL (read eval print loop)
func runRepl() {
	for {
		src, err := getInput()
		if err != nil || src == "exit" {
			return
		}

		run(src)
	}
}

// Prompts for input from the terminal
func getInput() (string, error) {
	bold := promptui.Styler(promptui.FGBold)
	prompt := promptui.Prompt{
		Label: ">",
		Validate: func(src string) error {
			if src == "exit" {
				return nil
			}

			_, err := parser.NewParser(lexer.FromString(src)).ParseProgramFully()
			err = parser.FormatErrorWithParserLocation(err)
			return err
		},
		Templates: &promptui.PromptTemplates{
			Prompt:  fmt.Sprintf("%s {{ . | bold }} ", bold(promptui.IconInitial)),
			Valid:   fmt.Sprintf("%s {{ . | bold }} ", bold(promptui.IconGood)),
			Invalid: fmt.Sprintf("%s {{ . | bold }} ", bold(promptui.IconBad)),
			Success: fmt.Sprintf("{{ . | faint }} "),
		},
	}
	return prompt.Run()
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
