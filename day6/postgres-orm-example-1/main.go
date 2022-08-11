package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func main() {
	dbURL := "postgres://dbuser1:password@192.168.1.254:5432/sampledb"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Book{})
	book := Book{Id: 1, Title: "programming in c", Author: "Byron", Desc: "Easy to learn"}
	if result := db.Create(&book); result.Error != nil {
		log.Fatal("Error creating record")
	}
	fmt.Println("record created")
}
