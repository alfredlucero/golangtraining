package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// net/http package for HTTP clients and servers

func main() {
	// GET request to server with http.Get
	// uses http.DefaultClient object
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// Print first 5 lines of response body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
