package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	// Simulate long-running task
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Request processed successfully")
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		}
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("server started at 8010")
	log.Fatal(http.ListenAndServe(":8010", nil))
}
