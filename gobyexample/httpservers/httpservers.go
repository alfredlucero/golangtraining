package main

import (
	"fmt"
	"net/http"
)

// Handlers = object implementing http.Handler interface
// Common way to write handler is by using the http.HandlerFunc adapter on functions with appropriate signature
// Take http.ResponseWriter and http.Request as arguments
// repsonse writer used to fill in the HTTP response
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// Reading all HTTP request headers and echoing them into response body
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Register handers on server routes using http.HandleFunc; sets up default router and takes function as argument
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Use default router and start on port 8090
	http.ListenAndServe(":8090", nil)

	// go run httpservers.go &
	// curl localhost:8090/hello
	// curl localhost:8090/headers
}
