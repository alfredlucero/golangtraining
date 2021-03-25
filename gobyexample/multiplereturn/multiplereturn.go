package main

import "fmt"

// Returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Multiple assignment to get the multiple return values from the function
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
