package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it
// Can communicate completion to channel's receivers
func main() {
	// When no more jobs for the worker, we'll close the jobs channel
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Worker goroutine that repeatedly receives jobs with j, more := <- jobs
	// more will be false if jobs has been closed and all values in the channel have already been received
	// notify on done when we've worked all our jobs
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Sends 3 jobs to the worker over the jobs channel then closes it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Await the worker using the synchronization approach
	<-done
}
