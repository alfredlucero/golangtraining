package contexts

import (
	"context"
	"fmt"
	"net/http"
)

// The server is not explicitly responsible for cancellation as it simply passes context and relies on the downstream functions to respect any cancellations that may occur.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}

type Store interface {
	// We send the context through to the downstream Store and it handles the error coming from Store when it is canceled.
	Fetch(ctx context.Context) (string, error)
}

// Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
// When a Context is canceled, all Contexts derived from it are also canceled.
// Pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests.
// It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly.

// Don't use context.Values as it's an untyped map and create a coupling of map keys from one module to another. Don't pass values through here unless you're doing things like
// passing in things orthogonal to a request like a trace id. Context.Value should inform, not control.
