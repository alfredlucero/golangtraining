package main

import (
	"fmt"
	"os"
)

// String formatting in the printf tradition

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// Prints an instance of our point struct
	fmt.Printf("%v\n", p)

	// If the value is a struct, the %+v variant will include the struct's field names
	fmt.Printf("%+v\n", p)

	// Provides Go syntax representation of the value
	fmt.Printf("%#v\n", p)

	// Printing type of value
	fmt.Printf("%T\n", p)

	// Formatting booleans
	fmt.Printf("%t\n", true)

	// Formatting integers
	fmt.Printf("%d\n", 123)

	// Binary representation
	fmt.Printf("%b\n", 14)

	// Character corresponding to integer
	fmt.Printf("%c\n", 33)

	// Hex encoding
	fmt.Printf("%x\n", 456)

	// Floats
	fmt.Printf("%f\n", 78.9)

	// Scientific notation
	fmt.Printf("%e\n", 1234500000.0)
	fmt.Printf("%E\n", 1234500000.0)

	// Base-16
	fmt.Printf("%x\n", "hex this")

	// Pointers
	fmt.Printf("%p\n", &p)

	// Width and precision of resulting figure
	// Width = number after % in the verb
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// Width of floats
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// Left-justify with - flag
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Basic right-justified width
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// Left-justify
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Format and return a string without printing it anywhere with Sprintf
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Format + print to io.Writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
