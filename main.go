package main

import (
	"github.com/howden/cham/repl"
	"os"
)

func main() {
	repl.HandleCommandLine(os.Args)
}
