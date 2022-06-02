package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// connect
	db, err := sql.Open("sqlite3", "./data.sql") //sql, db, sqlite, sqlite3
	if err != nil {
		panic(err)
	}

	// DDL / DML , create table, create data

	// sqlCreateUser := `CREATE TABLE peoples (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	name TEXT,
	// 	email TEXT,
	// 	address varchar,
	// 	age int
	// );`

	// _, err = db.Exec(sqlCreateUser)
	// if err != nil {
	// 	panic(err)
	// }

	// sqlInsertUser := `INSERT INTO peoples
	// (name, email, address, age)
	// VALUES (?, ?, ?, ?);`

	// _, err = db.Exec(sqlInsertUser, "John", "john@gmail.com", "Jalan Raya", "20")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("success")

	//update, delete, query

	sqlSelectUser := `SELECT * FROM peoples;`

	rows, err := db.Query(sqlSelectUser)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		var address string
		var age int

		err = rows.Scan(&id, &name, &email, &address, &age)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, name, email, address, age)
	}
}
