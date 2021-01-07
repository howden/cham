package main

import (
	"fmt"
	"github.com/howden/cham/lexer"
	"os"
	"strings"
)

func textLexer() {
	if len(os.Args) < 2 {
		fmt.Println("./cham <program>")
		return
	}

	src := strings.Join(os.Args[1:], " ")
	tokens, err := lexer.FromString(src).RemainingTokens()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, tok := range tokens {
		fmt.Println(tok)
	}
}

func main() {

}
