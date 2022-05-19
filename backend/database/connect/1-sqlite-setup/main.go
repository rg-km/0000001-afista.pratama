package main

import (
	"database/sql" // import direct
	"fmt"

	_ "github.com/mattn/go-sqlite3" // third party API , import indirect
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.sqlite")

	if err != nil {
		fmt.Println("error opening database:", err)
	}

	fmt.Println("You are successfully opening the database.")
	defer db.Close()
}

// melakukan koneksi ke sql kalian, RDBMS sqlite3

// input / output
// http => rest API, http server
// json => ".json" , example : 'data.json'
// io,
// file => .txt, .csv, .json
// database => sqlite3

// extensi => .db .sql .sqlite .sqlite3

// streaming, connection
