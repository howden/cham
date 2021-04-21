package repl

import (
	"fmt"
	"github.com/howden/cham/eval"
	"github.com/howden/cham/parser"
	"github.com/manifoldco/promptui"
)

// Runs the REPL (read eval print loop)
func StartRepl() {
	fmt.Println("CHAM Interpreter v1.0")
	store := eval.NewReactionStore()

	for {
		input, err := getInput(store)
		if err != nil {
			return
		}

		if len(input) > 0 && input[0] == ':' {
			command := input[1:]
			if command == "q" || command == "quit" {
				fmt.Println("Goodbye!")
				return
			} else if command == "s" || command == "store" {
				definitions := store.Slice()
				fmt.Printf("Showing all stored reactions: (%v)\n", len(definitions))
				for _, def := range definitions {
					fmt.Printf("- :%v\n", def.Identifier.Name())
				}
			} else {
				fmt.Printf("unknown command: %s\n", command)
			}
		} else {
			HandleReplInput(input, store)
		}
	}
}

func validateInput(input string, store *eval.ReactionStore) error {
	// command
	if len(input) > 0 && input[0] == ':' {
		command := input[1:]
		if command == "q" || command == "quit" || command == "s" || command == "store" {
			return nil
		}
		return fmt.Errorf("unknown command: %s", command)
	}

	// program
	_, _, err := ParseProgramOrReaction(input, store)
	err = parser.FormatErrorWithParserLocation(err)
	return err
}

// Prompts for input from the terminal
func getInput(store *eval.ReactionStore) (string, error) {
	bold := promptui.Styler(promptui.FGBold)
	prompt := promptui.Prompt{
		Label: ">",
		Validate: func(input string) error {
			return validateInput(input, store)
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
