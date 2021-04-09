package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Alfred's score", func(t *testing.T) {
		// Request method, path, body.
		request, _ := http.NewRequest(http.MethodGet, "/players/Alfred", nil)
		// Inspect what has been written as a response through ResponseRecorder.
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
