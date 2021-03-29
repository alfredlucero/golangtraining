package main

import "fmt"

// Accepts a channel for sending values, would be compile-time error to try to receive on this channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Accepts one channel for receives (pings) and second for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
