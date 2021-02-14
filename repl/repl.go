package repl

import (
	"fmt"
	"github.com/howden/cham/parser"
	"github.com/manifoldco/promptui"
)

// Runs the REPL (read eval print loop)
func StartRepl() {
	fmt.Println("CHAM Interpreter v0.1")
	for {
		input, err := getInput()
		if err != nil {
			return
		}

		if len(input) > 0 && input[0] == ':' {
			command := input[1:]
			if command == "q" || command == "quit" {
				fmt.Println("Goodbye!")
				return
			}
		}

		PrintEvalOutput(input)
	}
}

func validateInput(input string) error {
	// command
	if len(input) > 0 && input[0] == ':' {
		command := input[1:]
		if command == "q" || command == "quit" {
			return nil
		}
	}

	// program
	_, err := ParseProgram(input)
	err = parser.FormatErrorWithParserLocation(err)
	return err
}

// Prompts for input from the terminal
func getInput() (string, error) {
	bold := promptui.Styler(promptui.FGBold)
	prompt := promptui.Prompt{
		Label:    ">",
		Validate: validateInput,
		Templates: &promptui.PromptTemplates{
			Prompt:  fmt.Sprintf("%s {{ . | bold }} ", bold(promptui.IconInitial)),
			Valid:   fmt.Sprintf("%s {{ . | bold }} ", bold(promptui.IconGood)),
			Invalid: fmt.Sprintf("%s {{ . | bold }} ", bold(promptui.IconBad)),
			Success: fmt.Sprintf("{{ . | faint }} "),
		},
	}
	return prompt.Run()
}
