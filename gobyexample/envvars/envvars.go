package main

import (
	"fmt"
	"os"
	"strings"
)

// Environment variables for conveying configuration information to Unix programs

func main() {
	// Set a key/value pair with os.Setenv
	os.Setenv("FOO", "1")

	// Get a value for a key with os.Getenv
	fmt.Println("FOO: ", os.Getenv("FOO"))
	// Returns empty string if key isn't present in environment
	fmt.Println("BOO:", os.Getenv("BAR"))

	fmt.Println()
	// os.Environ to list all the key/value pairs in environment
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

}
