package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines
// Can send values into channels from one goroutine and receive those values into another goroutine
func main() {
	// make(chan val-type) - typed by values they convey
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax
	// Sending "ping" to messages channel from a new goroutine
	go func() { messages <- "ping" }()

	// <-channel syntax receives a value from the channel
	msg := <-messages
	fmt.Println(msg)

	// By default sends and receives block until both the sender and receiver are ready to allow us to wait at the end of our program for the "ping" message without having to use any other synchronization
}
