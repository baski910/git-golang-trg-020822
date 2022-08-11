package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=192.168.1.254 port=5432 user=dbuser1 password=password dbname=sampledb")
	if err != nil {
		log.Fatal("error connecting to db")
	}
	defer db.Close()
	/*
		err = db.Ping()
		if err != nil {
			log.Fatal("error connecting to db")
		}
		fmt.Println("connected to database")
	*/
	/*
		insertStmt := `insert into "t1"("id","name") values(1,'bob')`
		_, e := db.Exec(insertStmt)
	*/
	/*
		insertStmt := `insert into "t1"("id","name") values($1,$2)`
		_, e := db.Exec(insertStmt, 2, "tom")

		if e != nil {
			log.Fatal("error inserting record")
		}
		fmt.Println("record inserted")
	*/
	rows, err := db.Query(`SELECT * from t1`)
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		err = rows.Scan(&id, &name)
		fmt.Println(name, id)
	}

}
