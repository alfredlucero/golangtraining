package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Line filter is a common type of program that reads input on stdin, processes it, and then prints some derived result to stdout
// grep and sed are common line filters

// Example that writes a capitalized version of all input text
func main() {
	// Wrapping unbuffered os.Stdin with buffered scanner gives us a Scan method that advanced the scanner to the next token which is the next line in the default scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Text returns the current token here the next line from input
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// Check for errors during Scan; end of file is expected and not reported by Scan as an error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
