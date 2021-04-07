package selector

import (
	"fmt"
	"net/http"
	"time"
)

// Why are we testing the speeds of the websites one after another when Go is great at concurrency? We should be able to check at same time.
// We don't care about exact response times of the requests, we just want to know which one came back first.
// We'll use 'select', channels, and goroutines to help us out.
// Solution 1:
// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// Select lets you wait on multiple channels. The first one to send a value "wins" and the case is executed.
	select {
	// myVar := <-ch is a blocking call as you're waiting for a value.
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// Timeout case to prevent slow requests.
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// We don't care the type sent to the channel, we only want to signal we are done and close the channel.
// chan struct{} is smallest data type available from a memory perspective so we get no allocation versus a bool as we are closing and not sending anything.
func ping(url string) chan struct{} {
	// Always make channels as the zero value for channels is nil and if you try to send to it with <- it will block forever cause you cannot send to nil channels.
	ch := make(chan struct{})
	// Goroutine will send a signal into the channel once we have completed http.Get(url)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
