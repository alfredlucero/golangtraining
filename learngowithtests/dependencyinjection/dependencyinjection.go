package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
)

// We should be able to inject the dependency for printing.
// Our function doesn't need to care where or how the printing happens, so we should accept an interface rather than a concrete type.
// We can change the implementation of print to something we control so we can test it.
// For mocking, you replace real things you inject with a pretend version that you can control and inspect in your tests.

// Using io.Writer as our general purpose writer interface can be used for *bytes.Buffer and os.Stdout.
// We want to control where the data was written by injecting a dependency which allowed us to test our code, separate concerns, and allow code to be re-used in different contexts.
// i.e. Greet(os.Stdoout, "Alfred")
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// http.ResponseWriter works too as it implements io.Writer.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
