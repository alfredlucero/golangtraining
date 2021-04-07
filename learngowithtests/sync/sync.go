package sync

type Counter struct {
	// Mutex is a mutual exclusion lock. It must not be copied after first use.
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// All other goroutines will have to wait for this to unlock before proceeding.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

// Use channels when passing ownership of data.
// Use mutexes for managing state.
