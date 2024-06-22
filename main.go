package main

import (
	"fmt"
	"log"
	"net/http"
	"ascii-web/server"
)

func main() {
	// Register HomeHandler from server package to handle requests to "/"
	http.HandleFunc("/", server.HomeHandler)

	// Register Submit handler from server package to handle POST requests to "/ascii-art"
	http.HandleFunc("/ascii-art", server.Submit)

	// Print server URL to console
	fmt.Printf("Server running at http://localhost:8000/\n")

	// Start HTTP server on port 8000
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
