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
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
}

//var dbURL = "host=localhost user=demodbuser password=password dbname=demodb port=5432 sslmode=disable"
var dbURL = "postgres://dbuser2:password@192.168.1.254:5432/sampledb2"

func Router() *mux.Router {
	r := mux.NewRouter()
	return r
}

func main() {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Person{})
	mux := Router()
	mux.HandleFunc("/entries", GetEntries).Methods("GET")
	mux.HandleFunc("/entry/{id}", GetEntryById).Methods("GET")
	mux.HandleFunc("/entry", CreateEntry).Methods("POST")
	http.ListenAndServe(":8000", mux)
}

func GetEntries(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	//defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error connecting to db")
	}
	var persons []Person
	if result := db.Find(&persons); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "error fetching records")
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

func GetEntryById(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error connecting to db")
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var person Person
	if result := db.First(&person, id); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "error fetching record")
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)

}

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error connecting to db")
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	var person Person
	json.Unmarshal(body, &person)

	if result := db.Create(&person); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating record")
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Entry added")
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
