package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Use range to iterate over values received from a channel
	for elem := range queue {
		fmt.Println(elem)
	}
}
