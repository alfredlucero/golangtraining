package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Can also use built-in synchronization features of goroutines and channels to achieve same
// synchronization of state with explicit locking of mutexes
// Channel-based approach aligns with Go's ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine

// State will be owned by a single goroutine to ensure data is never corrupted by concurrent access
// To read/write state, other goroutines will send messages to the owning goroutine and receive corresponding replies
// These encapsulate those requests and a way for the goroutine to respond
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	// reads/writes channels used by other goroutines to issue read and write requests
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		// This goroutine owns the state in a private way
		var state = make(map[int]int)

		// Selects on the reads and writes channels, responding to requests as they arrive
		// Response is executed by first performing the requested operation and then sending a value on the response channel resp to indicate success and desired value in case of reads
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 100 goroutines to issue reads to the state-owning goroutine via the reads channel
	// Each read requires constructing a readOp, sending it over to the reads channel, and receiving the result over the provided resp channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10 writes with similar approach
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// This approach may be better when dealing with more channels or managing multiple mutexes are error-prone
}
