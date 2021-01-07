package main

import (
	"fmt"
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

	textLexer(src)
	testParser(src)
}

func textLexer(src string) {
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

func testParser(src string) {
	fmt.Println("Parser Output:")
	ast, err := parser.NewParser(lexer.FromString(src)).ParseProgramFully()
	if err != nil {
		if pe, ok := err.(*parser.ParserError); ok {
			fmt.Println(err)
			fmt.Printf("\n%s\n", src)
			fmt.Printf("%s^ HERE\n", strings.Repeat(" ", pe.LexerCurrentColumn-2))
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(ast)
	}
}
