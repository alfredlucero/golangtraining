package iteration

import (
	"strings"
)

func Repeat(character string, numRepeated int) string {
	// Old iteration way
	// var repeated string
	// for i := 0; i < numRepeated; i++ {
	// 	repeated += character
	// }
	// return repeated

	// Using strings library
	repeated := strings.Repeat(character, numRepeated)
	return repeated
}
