package main

import "fmt"

func main() {
	// Slices give a more powerful interface to sequences than arrays
	// Typed only by the elements they contain, not the number of elements
	// Build a slice with make
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// Append returns a slice containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f")

	// Slices can be copied with same length and contents
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice[low:high]
	l := s[2:5] // 2,3,4 elements
	fmt.Println("sl1:", l)

	l = s[:5] // up to but excluding 5
	fmt.Println("sl2:", l)

	l = s[2:] // up from and including 2
	fmt.Println("sl3:", l)

	// Declare and initialize a variable for slice in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Multi-dimensional data; length of inner slices can vary unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
