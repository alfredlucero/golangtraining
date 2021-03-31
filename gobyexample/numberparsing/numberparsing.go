package main

import (
	"fmt"
	"strconv"
)

// strconv package provides number parsing
func main() {
	// Parse float with 64 bits of precision
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Parse int and 0 means infer base from the string, 64 means result fits in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	// Recognizes hex-formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Base-10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return error on bad input
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
