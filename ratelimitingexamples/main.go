package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// Message struct for JSON responses
type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

// rateLimiter is an HTTP middleware that limits requests
func rateLimiter(next http.Handler) http.Handler {
	// Create a new limiter with a rate of 2 events per second and a burst of 4
	// This means it allows 2 requests per second and can handle bursts of up to 4 requests
	limiter := rate.NewLimiter(2, 4)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the request is not allowed by the limiter, return a 429 Too Many Requests status
		if !limiter.Allow() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			jsonResponse := `{"status": "Request Failed", "body": "The API is at capacity, try again later."}`
			_, err := w.Write([]byte(jsonResponse))
			if err != nil {
				log.Printf("Error writing response: %v", err)
			}
			return
		}

		// If allowed, serve the next handler
		next.ServeHTTP(w, r)
	})
}

// handler is the actual handler for the /message route
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse := `{"status": "Success", "body": "You accessed the message!"}`
	_, err := w.Write([]byte(jsonResponse))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	// Apply the rateLimiter middleware to the /message route
	http.Handle("/message", rateLimiter(http.HandlerFunc(handler)))

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
