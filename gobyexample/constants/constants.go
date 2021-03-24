package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5000000000

	const d = 3e20 / n
	fmt.Println(d)

	// Numeric constant has not type until it's given one such as by explicit conversion
	fmt.Println(int64(d))

	// Number given type float64 since funtion call required it
	fmt.Println(math.Sin(n))
}
