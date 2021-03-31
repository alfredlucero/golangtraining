package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Creating temporary file with ioutil.TempFile
	// Creates a file and opens it for reading and writing
	// "" for creating file in default location in OS
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// Unix-based OSes directory will likely be /tmp
	fmt.Println("Temp file name:", f.Name())

	// Clean up file after we're done even if OS will clean up files after some time
	defer os.Remove(f.Name())

	// Write some data to file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Create temporary directory and returns name rather than an open file
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
