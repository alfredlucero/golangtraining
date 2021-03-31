package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Dump a string/bytes into a file
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat", d1, 0644)
	check(err)

	// More granular writes can open a file for writing
	f, err := os.Create("/tmp/dat2")
	check(err)

	// Idiomatic to defer a Close immediately after opening a file
	defer f.Close()

	// Write byte slices as you'd expect
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Sync to flush writes to stable storage
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the underlying writer
	w.Flush()
}
