package main

import "fmt"

// Gets a copy by value for the function call
func zeroval(ival int) {
	ival = 0
}

// Pointers to pass references to values and records
func zeroptr(iptr *int) {
	// *iptr dereferences the pointer from its memory address to the current value at the address
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i gives memory address of i i.e. pointer to i
	// Changes value of i in main since it has a reference to the memory address for that variable
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too
	fmt.Println("pointer:", &i)
}
