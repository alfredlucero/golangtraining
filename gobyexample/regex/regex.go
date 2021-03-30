package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// Tests whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Compile an optimized Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods on these structs
	fmt.Println(r.MatchString("peach"))

	// Finds match for regexp
	fmt.Println(r.FindString("peach punch"))

	// Finds first match but returns start and end indices for the match instead of matching text
	fmt.Println(r.FindStringIndex("peach punch"))

	// Includes info about whole-pattern matches and submatches within those matches
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Apply to all matches in input
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Limit number of matches if provided non-negative integer
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Can drop String from function name with []byte arguments
	fmt.Println(r.Match([]byte("peach")))

	// When creating global variables with regex, can use the MustCompile variation of Compile
	// Panics instead of returning an error which makes it safer to use for global variables
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Can replace subsets of strings with other values
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func variant allows youo transform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
