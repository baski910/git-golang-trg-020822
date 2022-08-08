package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // multiplexing the connections
	mux.HandleFunc("/", home)
	log.Print("Starting on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
