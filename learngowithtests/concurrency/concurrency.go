package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// Concurrency means having more than one thing in progress.
// An operation that does not block in Go will run in a separate process called a goroutine.

// We use dependency injection with the WebsiteChecker to allow us to test the function without making HTTP calls.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Each loop will start a new goroutine and concurrent with the current process.
	// We need to make sure each result is added to the results map.
	for _, url := range urls {
		// Each goroutine has a reference to the url variable and not their own independent copy (it would write the url at the end of the iteration if we're not careful).
		// Thats why we pass in the url in each iteration to the goroutine anonymous function.
		go func(u string) {
			// We send a result struct for each call to wc to the resultChannel with a send statement (channel <- dataToSend).
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// There may be race conditions with concurrent map writes if we're not careful.
	// Race condition is a bug that occurs when software is dependent on the timing and sequence of events that we have no control over.
	// go test -race serves as our race detector.
	// We can solve this data race by coordinating our goroutines using channels.
	// Channels can both receive and send values to allow communication between different processes.
	// By sending results into a channel, we can control the timing of each write into the results map, ensuring it happens one at a time.
	for i := 0; i < len(urls); i++ {
		// Receive expression which assigns a value received from a channel to a variable (receivingVariable <- channel).
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
