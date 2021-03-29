package main

import (
	"fmt"
	"time"
)

// Can use channels to synchronize execution across goroutines
// Blocking receive to wait for a goroutine to finish
// Waiting for multiple goroutines to finish, you may use a WaitGroup

// Done channel used to notify another goroutine that his function's work is done
func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done
	done <- true
}

func main() {
	// Start worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel
	// If we removed the <- done line from the program, the program would exit before the worker even started
	<-done
}
