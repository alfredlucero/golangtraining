// Every Go program must be included in a package
// A standalone executable Go program must have a main function and must be included in the main package
// main is the entrypoint of execution
package main

import "fmt"

// go run helloworld.go -> runs main function and prints Hello world!
// go build helloworld.go -> builds executable you can run like ./helloworld.go
// go install helloworld.go -> installs into bin directory of Go workspace to be accessed globally like helloworld
func main() {
	fmt.Println("Hello world!")
}
