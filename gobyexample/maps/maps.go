package main

import "fmt"

func main() {
	// Maps are Go's built-in associative data type aka hashes or dicts
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key/value pairs
	m["k1"] = 7
	m["k2"] = 13

	// Print out all key-value pairs
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// Remove key/value pair
	delete(m, "k2")
	fmt.Println("map:", m)

	// Optional second return value when getting a value from a map indicates if the key was present in the map
	// Can disambiguate between missing keys and keys with empty values like 0 or ""
	_, present := m["k2"]
	fmt.Println("prs:", present)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
