package main

import (
	"fmt"
	"time"
)

// Rate limiting is useful for controlling resource utilization and maintaining quality of service
// Supported through goroutines, channels, and tickers
func main() {
	// Want to limit handling of incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Receive a value every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	// Blocking on receive from limiter channel before serving each request we can limit to 1 request ever 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Can buffer the limiter channel to allow bursts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200ms we try to add a new value to burstyLimiter up to its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Simulate 5 more incoming requests
	// First 3 benefit from burst capability
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
