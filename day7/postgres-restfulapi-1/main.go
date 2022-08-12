/*
product - struct
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
}

var dbURL = "postgres://dbuser1:password@192.168.1.254/sampledb"

func main() {

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Person{})
	mux := mux.NewRouter()
	mux.HandleFunc("/entries", GetEntries).Methods("GET")
	mux.HandleFunc("/entry/{id}", GetEntryById).Methods("GET")
	mux.HandleFunc("/entry", CreateEntry).Methods("POST")
	mux.HandleFunc("/entry/{id}", DeleteEntry).Methods("DELETE")
	http.ListenAndServe(":8000", mux)
}

func GetEntries(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error connecting to db")
	}
	var persons []Person

	if result := db.Find(&persons); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "error fetching records")
	}
	// db.First(&person,id)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	var person Person
	json.Unmarshal(body, &person)

	if result := db.Create(&person); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating records")
	}
}

func DeleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// find the person by id
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	var person Person

	if result := db.First(&person, id); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "error deleting records")
	}

	// Delete that person
	db.Delete(&person)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}

func GetEntryById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// Find person by Id
	var person Person

	if result := db.First(&person, id); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "error fetching records")
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Called for responses to encode and send json data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//encode payload to json
	response, _ := json.Marshal(payload)

	// set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
