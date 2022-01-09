package model

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func Init() {
// 	db, err := sql.Open("sqlite3", "../minierp.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	data, _ := db.Prepare("SELECT * FROM client;")
// 	dd, _ := data.Exec()
// 	fmt.Println(dd)
// 	defer db.Close()
// }
