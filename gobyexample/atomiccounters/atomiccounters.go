package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Primary mechanism for managing state in Go is communication over channels
// sync/atomic for atomic counters accessed by multiple goroutines

func main() {
	// Unsigned integer to represent always positive counter
	var ops uint64

	// WaitGroup to help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// Start 50 goroutines that each increment the counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			// Atomically increment the counter; provide the memory address of our ops counter
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Wait until all goroutines are done
	wg.Wait()

	fmt.Println("ops:", ops)

	// Without atomic ops, different numbers between runs and data race failures with --race flag would occur
}
