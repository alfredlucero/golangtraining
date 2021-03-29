package main

import "fmt"

// By default channels are unbuffered, meaning they will only accept sends (chan <-)
// if there is a corresponding receive (<- chan) ready to receive the sent value
// Buffered channels accept a limited number of values without a corresponding receiver for those values
func main() {
	// Channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// Can send values to channel without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// Receive these two values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
