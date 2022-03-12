package main

import (
	"database/sql"
	"log"
)

func dbConnection() {
	// Connect to database
	db, err = sql.Open("sqlite3", "./minierp.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	log.Println("db connection successful")
}
