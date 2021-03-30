package main

import (
	"fmt"
	"time"
)

// Worker will receive work on the jobs channel and send the corresponding results
// Sleep a second per job to simulate an expensive task
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// With pool of workers, we send work and collect results
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start up 3 workers initially blocked because there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs and then close that channel to indicate that's all the work we have
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect all the results of the work and ensures worker goroutines have finished; can use WaitGroup instead if using multiple goroutines
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	// 5 jobs executed by various workers concurrently ~ takes 2s for 5s of work
}
