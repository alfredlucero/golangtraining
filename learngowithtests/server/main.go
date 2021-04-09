package main

import (
	"log"
	"net/http"

	"example.com/playerserver/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Println("Started up server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
