package main

import "fmt"

// Takes two ints and returns sum as int
func plus(a int, b int) int {
	// Requires explicit returns
	return a + b
}

// Multiple consecutive parameters of the same type, you may omit the type name up to the final parameter that declares the type
func plusPlus(a, b, c int) int {
	return a + b + b
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
