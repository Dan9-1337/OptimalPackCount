package main

import (
	"OptimalPackCount/handler"
	"net/http"
)

func main() {
	// Set up the route for calculating packs.
	http.HandleFunc("/calculate-packs", handler.CalculatePacksHandler)

	// Serve static files from the "ui" directory.
	http.Handle("/", http.FileServer(http.Dir("./ui")))

	// Start the server on port 8080.
	http.ListenAndServe(":8080", nil)

}
