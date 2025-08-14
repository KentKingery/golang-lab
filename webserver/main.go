package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var logger = UTCLogger(false)

func main() {
	logger.Info("App startup")
	message := os.Getenv("WEB_MESSAGE")

	if message == "" {
		message = "No message defined"
	}

	// Define a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message) // Write the response to the client
	})

	// Start the HTTP server and listen on port 8080
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Listen and serve, logging any fatal errors
}
