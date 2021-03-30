package main

import (
	"fmt"
	"sort"
)

// sort package implements sorting for builtins and user-defined types
func main() {
	// Sort methods are specific to builtin type
	// Sorting is in-place so it changes the given slice and doesn't return a new one
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	// Can use sort to check if a slice is already in sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
