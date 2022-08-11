package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+") // abc@example.com

type Message struct {
	Email   string
	Content string
	Errors  map[string]string
}

// method receiver for validating messages
func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "Please enter a message"
	}

	return len(msg.Errors) == 0
}

func main() {
	mux := mux.NewRouter()
	/*
		mux.Get("/", http.HandlerFunc(home))
		mux.Post("/", http.HandlerFunc(send))
		mux.Get("/confirmation", http.HandlerFunc(confirmation))
	*/
	mux.HandleFunc("/", home).Methods("GET")
	mux.HandleFunc("/", send).Methods("POST")
	mux.HandleFunc("/confirmation", confirmation)
	log.Print("Listening...")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/home.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	// Step 1: Validate form
	msg := &Message{
		Email:   r.PostFormValue("email"),
		Content: r.PostFormValue("content"),
	}

	if msg.Validate() == false {
		render(w, "templates/home.html", msg)
		return
	}

	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
	// Step 2: Send message in an email
	// Step 3: Redirect to confirmation page
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
