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

	for i, dat := range data {
		if i != 0 {
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

	if err := u.LogoutAll(); err != nil {
		return nil, err
	}

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

	return nil, errors.New("Login Failed")
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

	return nil, errors.New("no user is logged in")
}

func (u *UserRepository) Logout(username string) error {
	//return nil

	// kita cari dulu user yang login, kemudian kita ubah loggedin menjadi false
	userLogin, err := u.FindLoggedinUser()
	if err != nil {
		return err
	}

	// if userLogin == &username {
	// 	u.changeStatus(username, false)
	// 	return nil
	// }

	// return errors.New("user not logged in")

	return u.changeStatus(*userLogin, false)
}

func (u *UserRepository) Save(users []User) error {
	record := [][]string{{
		"username", "password", "loggedin",
	}}

	for _, user := range users {
		newRow := []string{
			user.Username, user.Password,
		}

		if user.Loggedin {
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

	// error disini
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			users[i].Loggedin = status
			return u.Save(users)
		}
	}

	return errors.New("user not found")
}

func (u *UserRepository) LogoutAll() error {
	//return nil // TODO: replace this

	// dapatkan semua user, dan ubah semuanya menjadi false

	users, err := u.SelectAll()
	if err != nil {
		return err
	}

	var updateUsers []User

	for i := 0; i < len(users); i++ {
		if users[i].Loggedin {
			users[i].Loggedin = false
		}

		updateUsers = append(updateUsers, users[i])
	}

	err = u.Save(users)
	if err != nil {
		return err
	}

	return nil
}
