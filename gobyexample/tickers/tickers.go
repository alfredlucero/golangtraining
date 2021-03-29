package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future
// Tickers are for when you want to do something repeatedly at regular intervals
func main() {
	// Channel that is sent values like timers
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// Once a ticker is stopped, it won't receive any more values on its channel; we'll stop after 1600ms
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
