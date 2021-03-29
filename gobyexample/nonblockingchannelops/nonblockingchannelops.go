package main

import "fmt"

// Basic sends and receives on channels are blocking
// We can use select with a default clause to implement non-blocking sends, receives, and non-blocking multi-way selects

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive; if nothing in messages, it will immediately take the default case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Non-blocking send; msg cannot be sent to the messages channel because channel has no buffer and there is no receiver
	// Default is selected
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("No message sent")
	}

	// Multi-way non-blocking select
	// Non-blocking receives on both messages and signals
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
