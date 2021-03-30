package main

import "os"

func main() {
	// Panic = something went unexpectedly wrong
	// Fail fast on errors that shouldn't happen or we aren't prepared to handle gracefully
	panic("a problem at the disco")

	// Common use of panic is to abort if a function returns an error value that we don't know how to or want to handle
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	// Idiomatic in Go to use error-indicating return values wherever possible rather than exceptions
}
