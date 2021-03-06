package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Reading files requires checking most calls for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Slurping file's entire contents into memory
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// For more control over how and what parts of a file are read, can open the file to obtain an os.File value
	f, err := os.Open("/tmp/dat")
	check(err)

	// Read some bytes from the beginning of the file
	// Allow up to 5 to be read but also note how many were actually read
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek to a known location in the file and Read from there
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d btes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// io package provides functions for file reading
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind but Seek(0,0) does that
	_, err = f.Seek(0, 0)
	check(err)

	// bufio package implements a buffered reader that may be useful for its efficiency with many small reads and because of the additional reading methods it provides
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when you're done (usually this would be scheduled immediately after opening with defer)
	f.Close()
}
