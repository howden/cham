package repl

import (
	"fmt"
	"github.com/howden/cham/eval"
	"strings"
)

func HandleCommandLine(args []string) {
	if len(args) < 2 {
		StartRepl()
	} else if args[1] == "-h" || args[1] == "-help" || args[1] == "-version" {
		PrintHelp()
	} else if args[1] == "-f" || args[1] == "-file" {
		if len(args) < 3 {
			PrintHelp()
		} else {
			path := args[2]
			lines, err := readLines(path)
			if err != nil {
				fmt.Printf("error reading from file: %s\n", err)
				return
			}
			HandleFileInput(lines, eval.NewReactionStore())
		}
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
		HandleCmdLineInput(program)
	}
}

func PrintHelp() {
	fmt.Print(`CHAM Interpreter v1.0

  COMMANDS
    cham               Starts the REPL
    cham -h            Prints the help menu
    cham '<prog>'      Runs the given program and prints the output
    cham -f <file>     Loads and runs a program from the given file and
                       prints the output
    cham -l '<prog>'   Runs the given program through the lexer and prints
                       the output
    cham -p '<prog>'   Runs the given program through the parser and prints
                       the output

  REPL USAGE
    Enter a program into the prompt, then press enter to evaluate it.
    A red cross is displayed at the prompt for invalid input. A green tick is
    displayed for valid input.

    If there was a program parsing the program, an error message will be displayed.

  REPL COMMANDS
    :quit   :q    quit the REPL
    :load   :l    loads programs from the given file (provided as an argument)
    :store  :s    view a list of reactions saved in the REPLs memory

`)
}
