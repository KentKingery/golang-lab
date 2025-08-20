package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.Info("Dummy service started")
	message := os.Getenv("WEB_MESSAGE")

	if message == "" {
		message = "No message defined"
	}

	webpage := "<html><head><title>Dummy Service</title></head><body><h1>" + message + "</h1></body></html>"

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	// Define a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Writing page")
		slog.Info(webpage)
		fmt.Fprintf(w, webpage) // Write the response to the client
	})

	// Start the HTTP server and listen on port 8080
	fmt.Println("Server starting on port 8080...")
	slog.Error(http.ListenAndServe(":8080", nil).Error()) // Listen and serve, logging any fatal errors

	slog.Info("Dummy service stopped")
}
