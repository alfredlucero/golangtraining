package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alfred")

	got := buffer.String()
	want := "Hello, Alfred"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
