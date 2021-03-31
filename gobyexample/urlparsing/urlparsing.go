package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// Parsing URL which has a scheme, authentication info, host, port, path, query params and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)

	// User contains all the authentication info
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host contains both hostname and port if present
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// To get query params in a string of k=v format, use RawQuery
	// Can also parse query params into a map
	// Parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
