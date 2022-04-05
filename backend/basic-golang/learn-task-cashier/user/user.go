package user

import "fmt"

// golang , default huruf depan huruf kecil (private) ke folder / dir lain
type user struct {
	name    string
	age     int
	address string
	hobbies []string
}

func NewUser(name string, age int, address string, hobbies []string) user {
	return user{
		name:    name,
		age:     age,
		address: address,
		hobbies: hobbies,
	}
}

func (u user) GetName() string {
	return u.name
}

func (u user) String() string {
	return fmt.Sprintf("nama: %s, umur : %v, address: %v", u.name, u.age, u.address)
}
