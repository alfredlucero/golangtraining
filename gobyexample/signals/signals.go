package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Handling Unix signals i.e. server might want to gracefully shutdown when it receives a SIGTERM
// command-line tool to stop processing input if it receives a SIGINT

func main() {
	// sending os.Signal values on a channel sigs
	sigs := make(chan os.Signal, 1)
	// channel to notify us when the program can exit
	done := make(chan bool, 1)

	// Registers the given channel to receive notifications of specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// goroutine executes a blocking receive for signals
	// When it gets one it'll print it out and notify the program it can finish
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Program will wait here until it gets expected signal
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	// typing Ctrl-C aka ^C will send a SIGINT signal causing the program to print interrupt and then exit
}
