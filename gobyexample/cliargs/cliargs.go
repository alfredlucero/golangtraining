package main

import (
	"fmt"
	"os"
)

// Command-line arguments are a common way to parameterize execution of programs
// go run hello.go uses run and hello.go

func main() {
	// os.Args provides access to raw command-line arguments
	argsWithProg := os.Args
	// First value in slice is path to program; rest holds arguments to program
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
