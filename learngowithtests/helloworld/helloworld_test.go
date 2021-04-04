package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Alfred", "spanish")
		want := "Hola, Alfred"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Alfred", "french")
		want := "Bonjour, Alfred"
		assertCorrectMessage(t, got, want)
	})
}
