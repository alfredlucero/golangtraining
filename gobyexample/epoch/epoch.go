package main

import (
	"fmt"
	"time"
)

// Common to get number of seconds, milliseconds, or nanoseconds since the Unix epoch

func main() {
	// time.Now with Unix or UnixNano to get elapsed time since Unix epoch in seconds or nanoseconds
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// There is no UnixMillis so need to divide from nanos
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// Convert integer seconds or nanoseconds since the epoch into the corresponding time
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
