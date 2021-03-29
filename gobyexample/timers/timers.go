package main

import (
	"fmt"
	"time"
)

// Execute Go code at some point in the future or repeatedly at some interval
// timer and ticket features help us with this

func main() {
	// Timers represent a single event in the future
	// Tell timer how long you want to wait and provides a channel that will be notified at that time
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks on the timer's channel C until it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you wanted to wait you could do time.Sleep
	// Can also cancel the timer before it fires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
