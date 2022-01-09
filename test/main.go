package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./minierp.db")

	data, _ := db.Prepare(`
			INSERT INTO newsfeed(id, content) VALUES(?,?)
	`)
	data.Exec(2, "Abdullah")
}
