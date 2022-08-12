package main

import (
	"net/http"
	"postgres-restful-api-2/db"
	"postgres-restful-api-2/handlers"

	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods("GET")
	router.HandleFunc("/book", h.CreateBook).Methods("POST")
	http.ListenAndServe(":8000", router)
	//fmt.Println(h)
}
