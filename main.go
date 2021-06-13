package main

import (
	"ASMLang/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
