package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joaovds/chegae-server/internal/tracking"
)

func main() {
	mainMux := http.NewServeMux()

	tracking.NewModule().SetupHandlers(mainMux)

	log.Println("Server running on port", 8080)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), mainMux); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}
