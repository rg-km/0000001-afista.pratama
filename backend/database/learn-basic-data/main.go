package main

import "fmt"

// database
// table user, column : name, address, age
// table school, column : name, address

// database relationship
// 1. one to one
// 1 row table user, 1 row table school
// 2. one to many
// 1 data row table, berelasi / berhubungan dengan banyak row di table lain (bisa lebih dari 1)
// 3. many to many

type UserRow struct {
	ID             int // primary Key
	Name           string
	Address        string
	Age            int
	SchoolID       int // Foreign Key
	PhoneNumberIDs []int
}

type PhoneNumberRow struct {
	ID          int // primary Key
	CountryCode int
	Number      int
}

type SchoolRow struct {
	ID      int // primary Key
	Name    string
	Address string
}

type UserTable []UserRow
type SchoolTable []SchoolRow
type PhoneNumberTable []PhoneNumberRow

type Database struct {
	Users        UserTable
	Schools      SchoolTable
	PhoneNumbers PhoneNumberTable
}

func main() {
	var database Database

	var userRow1 = UserRow{1, "John", "Jakarta", 23, 1, []int{1, 2}}
	var userRow2 = UserRow{2, "Jane", "Semarang", 24, 2, []int{3}}
	var userRow3 = UserRow{3, "John", "Bandung", 24, 1, []int{4}}
	var userRow4 = UserRow{4, "Bernard", "Jakarta", 30, 1, []int{5}}

	var userTable = UserTable{userRow1, userRow2, userRow3, userRow4}

	database.Users = userTable

	var schoolRow1 = SchoolRow{1, "SMA N 1", "Jakarta"}
	var schoolRow2 = SchoolRow{2, "SMA N 2", "Semarang"}
	var schoolRow3 = SchoolRow{3, "SMA N 3", "Jakarta"}

	var schoolTable = SchoolTable{schoolRow1, schoolRow2, schoolRow3}

	database.Schools = schoolTable

	// one to many
	// i row table bisa berelasi lebih dari 1 data row di table lain, minimal ada 1 row
	var pnRow1 = PhoneNumberRow{1, 62, 123456789}
	var pnRow2 = PhoneNumberRow{2, 62, 234567890}
	var pnRow3 = PhoneNumberRow{3, 62, 345678901}
	var pnRow4 = PhoneNumberRow{4, 62, 456789012}
	var pnRow5 = PhoneNumberRow{5, 62, 567890123}

	var pnTable = PhoneNumberTable{pnRow1, pnRow2, pnRow3, pnRow4, pnRow5}

	database.PhoneNumbers = pnTable

	// Primary Key
	// 1. ID, identity, (unique identify)
	// memanfaatkan index si array
	// index dimulai dari 0
	// urutan / angka dimulai dari 1

	// row 2 , 2 - 1 ,  1
	// row 3 , 3 - 1 ,  2

	var id = 2

	for _, user := range database.Users {
		if user.ID == id {
			fmt.Println("User ID", id, ":", user)
		}
	}
}

// many to many
// 1 atau lebih data di row table punya relasi dengan 1 atau lebih data di row table lain

// user
// barang

// user1 barang1
// user1 barang2
// user2 barang1
// user2 barang3
// user3 barang2

// user1 => barang 1 dan barang 2
// user2 => barang 1 dan barang 3

// barang 1 => user 1 dan user 2
// barang 2 => user 1 dan user 3
