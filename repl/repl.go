package repl

import (
	"fmt"
	"github.com/howden/cham/eval"
	"github.com/howden/cham/parser"
	"github.com/manifoldco/promptui"
	"strings"
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

		isCommand, command, args := parseCommand(input)
		if isCommand {
			if command == "q" || command == "quit" {
				// quit command
				fmt.Println("Goodbye!")
				return

			} else if command == "l" || command == "load" {
				// load command
				if len(args) == 0 {
					fmt.Println("You need to specify a filename!")
					continue
				}

				path := args[0]
				lines, err := readLines(path)
				if err != nil {
					fmt.Printf("error reading from file: %s\n", err)
					continue
				}

				HandleFileInput(lines, store)
				fmt.Println("OK")

			} else if command == "s" || command == "store" {
				// store command
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
	isCommand, command, _ := parseCommand(input)
	if isCommand {
		if command == "q" || command == "quit" ||
			command == "s" || command == "store" ||
			command == "l" || command == "load" {
			return nil
		} else {
			return fmt.Errorf("unknown command: %s", command)
		}
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

// Attempts to parse a command from the given input
// Returns (whether a command was parsed, the command name, any extra arguments)
func parseCommand(input string) (bool, string, []string) {
	if len(input) == 0 {
		return false, "", nil
	}

	args := strings.Split(input, " ")
	if len(args) == 0 {
		return false, "", nil
	}

	command := args[0]
	if len(command) < 2 || command[0] != ':' {
		return false, "", nil
	}

	return true, command[1:], args[1:]
}
