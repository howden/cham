package repl

import (
	"fmt"
	"strings"
)

func HandleCommandLine(args []string) {
	if len(args) < 2 {
		StartRepl()
	} else if args[1] == "-h" || args[1] == "-help" || args[1] == "-version" {
		PrintHelp()
	} else if args[1] == "-l" {
		if len(args) < 3 {
			PrintHelp()
		} else {
			program := strings.Join(args[2:], " ")
			PrintLexerOutput(program)
		}
	} else if args[1] == "-p" {
		if len(args) < 3 {
			PrintHelp()
		} else {
			program := strings.Join(args[2:], " ")
			PrintParserOutput(program)
		}
	} else {
		program := strings.Join(args[1:], " ")
		PrintEvalOutput(program)
	}
}

func PrintHelp() {
	fmt.Print(`CHAM Interpreter v0.1

  COMMANDS
    cham               Runs the REPL
    cham -h            Prints the help menu
    cham '<prog>'      Runs the given program and prints the output
    cham -l '<prog>'   Runs the given program through the lexer and prints
                       the output
    cham -p '<prog>'   Runs the given program through the parser and prints
                       the output

  REPL USAGE
    Enter a program into the interpreter prompt, then hit enter to evaluate it.
    A red cross is displayed at the prompt for invalid input. A green tick is
    displayed for valid input.

    If there was a program parsing the program, an error message will be displayed.

  REPL COMMANDS
    :quit :q   quit the REPL

`)
}
