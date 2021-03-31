package main

import (
	"flag"
	"fmt"
)

// Flags to specify options for command-line programs

func main() {
	// Basic flag declarations for string, integer, boolean options
	// Returns string pointer
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string")

	// Once all flags declared, call flag.Parse() to execute the command-line parsing
	flag.Parse()

	// Dump out the parsed options and any trailing positional arguments
	// Need to dereference the points with *wordPtr to get the actual option values
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	// Values for all flags
	// $ ./command-line-flags -word=opt -numb=7 -fork -svar=flag

	// If you omit flags they automatically take their default values
	// $ ./command-line-flags -word=opt

	// Trailing positional arguments can be provided after any flags
	//	$ ./command-line-flags -word=opt a1 a2 a3

	// flag package requires all flags to appear before positional arugments (otherwise the flags will be interpreted as positional arguments)
	// $ ./command-line-flags -word=opt a1 a2 a3 -numb=7

	// Use -h or --help flags to get automatically generated help text for the command-line program
	// $ ./command-line-flags -h

	// Provide a flag that wasn't specified to the flag package, program will print an error message and show help text again
	// $ ./command-line-flags -wat
}
