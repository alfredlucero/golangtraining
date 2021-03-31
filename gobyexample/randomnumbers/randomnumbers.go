package main

import (
	"fmt"
	"math/rand"
	"time"
)

// math/rand for pseudorandom number generation

func main() {
	// Random int n where 0 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// Random float
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Default number generator is deterministic, so it'll produce the same sequence of numbers each time by default
	// To produce varying sequences, give it a seed that changes
	// Use crypto/rand for secret random numbers instead
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Call resulting rand.Rand just like functions on the rand package
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// If you seed a source with the same number, it produces the same sequence of numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
