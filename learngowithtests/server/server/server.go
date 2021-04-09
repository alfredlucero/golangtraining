package server

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	// ResponseWriter implements io.Writer so we can use fmt.Fprint to send strings as HTTP responses
	fmt.Fprint(w, "20")
}
