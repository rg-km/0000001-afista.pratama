package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

func (r *EmployeeRepository) CreateEmployee(employee *model.Employee) (int64, error) {
	var sqlStmt string

	// Task :
	// Membuat query untuk menambahkan data employee ke dalam database
	// Parameter : nik, first_name, last_name, email
	// Lihat model.Employee untuk data yang dibutuhkan

	// TODO: answer here
	sqlStmt = `INSERT INTO employees 
		(nik, first_name, last_name, email) 
		VALUES (?, ?, ?, ?);`

	// nik 1231231
	// john
	// doe
	// johndoe@gmail.com

	// INSERT INTO employees (nik, first_name, last_name, email) VALUES (123123, john, doe, johndoe@gmail.com)
	result, err := r.db.Exec(sqlStmt, employee.NIK, employee.FirstName, employee.LastName, employee.Email)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
