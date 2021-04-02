package main

import (
	"fmt"
	"net/http"
	"time"
)

// HTTP servers useful for demonstrating usage of context.Context for controlling cancellation
// Context carries deadlines, cancellation signal

func hello(w http.ResponseWriter, req *http.Request) {
	// context.Context created for each request by the net/http
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait for a few seconds before sending reply to client
	// While woring, keep an eye on the context's Done() channel for a signal that we should cancel the work and return as soon as possible
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// Returns error explaining why Done() channel was closed
		err := ctx.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)

	// go run context.go
	// curl localhost:8090/hello
	// ^C
}
