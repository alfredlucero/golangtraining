package main

import "fmt"

func main() {
	// Range iterates over elements in a variety of data structures

	// Using range to sum numbers in a slice; arrays work like this too
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Ranges on arrays and slices provide both the index and value for each entry
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Can iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range on strings iterates over Unicode code points
	// First value is starting byte index of the rune and the second the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
