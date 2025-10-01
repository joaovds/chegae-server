package main

import (
	"log"
	"net/http"

	trackingHttp "github.com/joaovds/chegae-server/internal/tracking/adapters/http"
)

func main() {
	http.HandleFunc("GET /ws", trackingHttp.WsHandler)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
