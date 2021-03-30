package main

// Standard library's strings package provides string-related functions
import (
	"fmt"
	s "strings"
)

// Alias fmt.Println to shorter name
var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace:", s.Replace("foo", "o", "0", 1))
	p("Split:", s.Split("a-b-c-d-e", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
	p()

	// Go uses UTF-8 encoded strings; need encoding aware operations if dealing with multi-byte characters
	// Getting length of string in bytes
	p("Len:", len("hello"))
	// Getting byte by index
	p("Char:", "hello"[1])

}
