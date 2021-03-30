package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Get current time
	now := time.Now()
	p(now)

	// Build a time struct by providing the year, month, day, etc.
	// Times are always associated with Location i.e. timezone
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Extract various time component values
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Monday - Sunday
	p(then.Weekday())

	// Comparing two times
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Returns a duration representing interval between two times
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Advance a time by a given duration or backwards by a duration
	p(then.Add(diff))
	p(then.Add(-diff))
}
