package main

import (
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Name       string
	Occupation string
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/layout.html"))

		data := []User{
			{Name: "John Doe", Occupation: "gardener"},
			{Name: "Roger Roe", Occupation: "driver"},
			{Name: "Thomas Green", Occupation: "teacher"},
		}
		tmpl.Execute(w, data)

	})

	log.Println("server started at 8010")
	log.Fatal(http.ListenAndServe(":8010", nil))
}
