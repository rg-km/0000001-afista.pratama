package repository

import (
	"errors"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	// tujuannya untuk membuka data dari db atau create filenya
	// TODO: replace this
	data, err := u.db.Load("users")
	if err != nil {
		var record = [][]string{{
			"username", "password", "loggedin",
		}}

		if err := u.db.Save("users", record); err != nil {
			return nil, err
		}
	}

	var users []User

	for _, dat := range data {
		user := User{
			Username: dat[0],
			Password: dat[1],
		}

		if dat[2] == "true" {
			user.Loggedin = true
		} else {
			user.Loggedin = false
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	// TODO: replace this
	return u.LoadOrCreate()
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	// TODO: replace this
	// jika username yang dimasukkan di terminal == yang ada di users.csv
	// dan password di terminal == password di users.csv
	// berarti dapat akses login

	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			u.changeStatus(username, true)
			return &username, nil
		}
	}

	return nil, errors.New("username & password invalid")
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	// TODO: replace this

	// mencari user yang loggedin nya true
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Loggedin {
			return &user.Username, nil
		}
	}

	return nil, errors.New("user not login")
}

func (u *UserRepository) Logout(username string) error {
	return nil // TODO: replace this
}

func (u *UserRepository) Save(users []User) error {
	record := [][]string{{
		"username", "password", "loggedin",
	}}

	for _, user := range users {
		newRow := []string{
			user.Username, user.Password,
		}

		if user.Loggedin == true {
			newRow = append(newRow, "true")
		} else {
			newRow = append(newRow, "false")
		}

		record = append(record, newRow)
	}

	return u.db.Save("users", record)
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			users[i].Loggedin = true
		}
	}

	return u.Save(users)
}

func (u *UserRepository) LogoutAll() error {
	return nil // TODO: replace this
}
