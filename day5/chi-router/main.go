package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func main() {
	r := chi.NewRouter()

	r.HandleFunc("/", Greet)
	log.Println("starting server at port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
