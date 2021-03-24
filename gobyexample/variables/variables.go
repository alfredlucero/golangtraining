package main

import "fmt"

func main() {
	// Variables explicitly declared and used by compiler to check type-correctness of function calls
	var a = "initial"
	fmt.Println(a)

	// Declare multiple variables at once
	var b, c int = 1,2
	fmt.Println(b,c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued i.e. int is 0
	var e int
	fmt.Println(e)

	// := syntax shorthand for declaring and initializing a variable
	f := "apple"
	fmt.Println(f)
}
