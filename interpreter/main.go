package main

import (
	"fmt"
	"os"
	"interpreter/repl"
)

func main() {
	fmt.Printf("Welcome to Omega!\n")
	repl.Start(os.Stdin, os.Stdout)
}
