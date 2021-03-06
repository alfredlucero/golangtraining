package main

import "fmt"

// Take arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// Can apply multiple args in a slice like func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
