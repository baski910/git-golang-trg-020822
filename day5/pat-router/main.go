package main

import (
	"fmt"
	"net/http"

	"github.com/bmizerany/pat"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func ReplyNamePat(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	name := parameters.Get(":name")
	w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
}

func main() {
	mx := pat.New()
	mx.Get("/", http.HandlerFunc(SayHelloWorld))
	mx.Get("/:name", http.HandlerFunc(ReplyNamePat))

	http.Handle("/", mx)
	http.ListenAndServe(":8000", nil)
}
