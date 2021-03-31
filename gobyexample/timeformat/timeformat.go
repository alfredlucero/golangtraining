package main

import (
	"fmt"
	"time"
)

// Time formatting and parsing via pattern-based layouts

func main() {
	p := fmt.Println

	// Format time to RFC3339
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Parsing uses same layout values as Format
	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	// Format and Parse use example-based layouts
	// Can supply custom layouts, must use reference time to show pattern
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// Purely numeric representations can use standard string formatting with extracted components of the time value
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// Parse returns an error on malformed input
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
