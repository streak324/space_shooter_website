package main

import (
	"log"
	"net/http"
)

func main() {
	// File server to serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))

	// Handle requests to the root URL by serving files from the "public" directory
	http.Handle("/", fs)

	// Define the port to listen on
	addr := "127.0.0.1:8080"

	// Start the web server
	log.Printf("Starting web server on http://%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
