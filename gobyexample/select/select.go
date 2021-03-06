package main

import (
	"fmt"
	"time"
)

// Go's select lets you wait on multiple channel operations
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	//  Each channel will receive a value after some amount of time to simulat blocking RPC operations executing in concurrent goroutines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// Select to await both of these values simultaneously and printing each one as it arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
