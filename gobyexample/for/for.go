// for is Go's only looping construct
package main

import "fmt"

func main() {
	i := 1

	// Single condition loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic initial/condition/after
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Without a condition will loop repeatedly until you break/return out of the loop
	for {
		fmt.Println("loop")
		break
	}

	// Continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
