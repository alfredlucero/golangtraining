package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// filepath provides functions to parse and construct file paths in a way that is portable between operating systems
// dir/file on Linux vs. dir\file on Windows

func main() {
	// Join should be used to construct paths in a portable way; constructs hierarchical path from them
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Use Join instead of concatenating manually; it will normalize paths by removing superfluous separators and directory changes
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Can split path to directory and file; Split will return both in the same call
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Check absolute path
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Split extension out of names with Ext
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Find file name with extension removed
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds relative path between a base and a target
	// Returns an error if target cannot be made relative to base
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
