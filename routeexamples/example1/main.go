package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewCon() *gorm.DB {
	dsn := "root:admin_12345@tcp(127.0.0.1:3306)/demodb_1"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("unable to connect to db")
	}
	db.AutoMigrate(&Product{})
	return db
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", GetMethod).Methods("GET")
	router.HandleFunc("/api/products", PostMethod).Methods("POST")
	log.Println("server started at 8010")
	log.Fatal(http.ListenAndServe(":8010", router))
}

func GetMethod(w http.ResponseWriter, r *http.Request) {
	db := NewCon()
	var products []Product
	result := db.Find(&products)
	if result.Error != nil {
		fmt.Println("unable to fetch products")
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Get method called")
}

func PostMethod(w http.ResponseWriter, r *http.Request) {
	r *http.Request) {
	// 1. connect to db
	db := NewCon()
	// 2. create varible of type product
	var product Product
	// 3. assign the values posted to the above variable
	json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close() // close the stream
	//4. create a new record using the posted value
	if result := db.Create(&product); result.Error != nil {
		fmt.Println("unable to create product")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
	fmt.Fprintf(w, "Post method called")
}
