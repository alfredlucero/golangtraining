package main

import (
	"fmt"
	"os"
)

// os.Exit to immediately exit with a given status

func main() {
	// defers will not be run when using os.Exit so this is never called
	defer fmt.Println("!")

	os.Exit(3)
}
