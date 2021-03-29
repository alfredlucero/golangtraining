package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Runs synchronously
	f("direct")

	// go f(s)
	// goroutine is a lightweight thread of execution
	// will execute concurrently with the calling one
	go f("goroutine")

	// can start goroutine for an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two functions are running asynchronously in separate goroutines for now
	// wait for them to finish using WaitGroup
	time.Sleep(time.Second)
	fmt.Println("done")

	// We should see blocking call first then the interleaved output of two goroutines, reflecting how the goroutines are
	// run concurrently by the Go runtime
}
