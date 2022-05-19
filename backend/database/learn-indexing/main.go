package main

import "fmt"

type UserRow struct {
	ID      int // primary Key
	Name    string
	Address string
	Age     int
}

type UserTable []UserRow

type Database struct {
	Users UserTable
}

// method, CRUD
// search / read, insert / create, update, delete

func (d *Database) Search(id int) *UserRow {

	// O(n), n akan terus meningkat sesuai banyaknya eksekusi
	for _, user := range d.Users {
		if user.ID == id {
			fmt.Println("user row found")
			return &user
		}
	}

	fmt.Println("error id not found")
	return nil
}

func (d *Database) Insert(user UserRow) {
	d.Users = append(d.Users, user)
}

func (d *Database) Update(id int, user UserRow) {
	for i, user := range d.Users {
		if user.ID == id {
			d.Users[i] = user
			fmt.Println("user row updated")
			return
		}
	}
}

// 1000
// id 900
// 1, 2, ,3 ,4 , ....., 900 == id , kita update

func (d *Database) Delete(id int) {
	for i, user := range d.Users {
		if user.ID == id {
			d.Users = append(d.Users[:i], d.Users[i+1:]...)
			fmt.Println("user row deleted")
			return
		}
	}
}

// hashmap, atau indexing biar lebih efisien
// UserRow

type PrimaryKey int

type UserTableMap map[PrimaryKey]UserRow

func (u *UserTableMap) Insert(user UserRow) {
	// len = 0 +1 , 1
	// O(1)
	(*u)[PrimaryKey(len(*u))+1] = user
}

func (u *UserTableMap) Search(id PrimaryKey) *UserRow {
	// O(1)
	data := (*u)[id]
	return &data
}

func (u *UserTableMap) Update(id PrimaryKey, Name string, Address string, age int) {
	// O(1)
	data := (*u)[id]

	data.Name = Name
	data.Address = Address
	data.Age = age

	(*u)[id] = data
}

func (u *UserTableMap) Delete(id PrimaryKey) {
	// delete
	// O(1)
	delete((*u), id)
}

func main() {
	// insert 1000 data

	var users []UserRow

	for i := 0; i < 1000; i++ {
		users = append(users, UserRow{i + 1, "John", "Jakarta", 24})
	}

	userTable := users

	databases := Database{
		Users: userTable,
	}

	databases.Search(999)
}
