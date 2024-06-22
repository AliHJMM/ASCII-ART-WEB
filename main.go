package main

import (
	"fmt"
	"log"
	"net/http"
	"ascii-web/server"
)

func main() {
	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/ascii-art", server.Submit)
	fmt.Printf("http://localhost:8000/\n")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
