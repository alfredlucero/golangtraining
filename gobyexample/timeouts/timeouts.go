package main

import (
	"fmt"
	"time"
)

// Timeouts important for things that connect to external resources or bounding execution time
// Can implement timeouts with channels and select
func main() {
	// Simulate non-blocking send to channel based on a fake external call
	// Buffered channel with size 1 to prevent goroutine leaks in case the channel is never read
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	// awaits a value to be send after the timeout of 1s
	// select proceeds with first receive that's ready so we'll take the timeout case if operation takes more than allowed 1s
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// Longer timeout of 3s, then the receive from c2 will succeed and we'll print the result
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
