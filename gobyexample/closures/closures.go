package main

import "fmt"

// Go supports anonymous functions which can form closures
// Returns another function which we anonymously define in the body
// Returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Assigns the resulting function that bumps the count
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// State is unique to the particular function
	newInts := intSeq()
	fmt.Println(newInts())
}
