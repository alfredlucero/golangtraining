package main

import (
	"fmt"
	"sort"
)

// In order to sort by a custom function in Go, we need a corresponding type
// This aliases the builtin []string type
type byLength []string

// We implement sort.Interface - Len, Less, and Swap on our type so we can use the sort package's generic Sort function
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// By following the same pattern of creating a custom type, implementing the three Interface methods on that type, and then calling sort.Sort on a collection of that custom type, we can sort Go slices by arbitrary functions
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// Convert original fruits slice to byLength and then use sort.Sort on the typed slice
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
