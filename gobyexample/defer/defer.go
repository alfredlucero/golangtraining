package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program's exuection, usually for purposes of cleanup
// defer used where ensure and finally would be used in other languages

func main() {
	f := createFile("/tmp/defer.txt")
	// Defer closing of file to be executed at end of enclosing function after writeFile has finished
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Important to check for errors when closing a file
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
