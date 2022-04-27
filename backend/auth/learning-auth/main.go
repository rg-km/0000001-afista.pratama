package main

import (
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// package HTTP => I/O
// write (output), request (input)

// authentication / authorization (security)
// encrypt sama decrypt
// JWT

// basic authorization

// - bahas pas workshop
// session key auth
// token base auth
// - JWT

const (
	USERNAME = "afista"
	PASSWORD = "afista1234"
)

var users = []*User{
	{Username: "afista", Password: "afista1234", Role: "admin"},
	{Username: "user2", Password: "user2", Role: "user"},
	{Username: "user3", Password: "user3", Role: "user"},
}

type User struct {
	Username string
	Password string
	Role     string
}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

var students = []*Student{
	{Id: "001", Name: "ary", Grade: 1},
	{Id: "002", Name: "ary 2", Grade: 2},
	{Id: "003", Name: "ary 3", Grade: 2},
}

func checkUser(username string, password string) (User, bool) {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return *u, true
		}
	}

	return User{}, false
}

func Auth(w http.ResponseWriter, r *http.Request) (string, bool) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("error happen"))
		return "", false
	}

	user, ok := checkUser(username, password)
	if !ok {
		w.Write([]byte("invalid username / password"))
		return "", false
	}

	return user.Role, true
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	// struct to json
	data, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// handlerfunction
// ActionStudent is function handler to send the student data
func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return
	}

	role, ok := Auth(w, r)
	if !ok {
		return
	}

	// authorization
	log.Println("basic auth")
	if role == "admin" {
		OutputJSON(w, students)
	} else {
		OutputJSON(w, struct {
			Message string
		}{
			Message: "only for admin role",
		})
	}
}

func main() {
	http.HandleFunc("/student", ActionStudent)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, ok := Auth(w, r)
		if !ok {
			return
		}

		// akses route yang dibatasi
		w.WriteHeader(200)
		w.Write([]byte("success GET root route / "))
	})

	fmt.Println("server running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func EncrypDecrypt() {
	var secret = "ini data yang sangat rahasia"

	key := base64.StdEncoding.EncodeToString([]byte(secret))
	// aW5pIGRhdGEgeWFuZyBzYW5nYXQgcmFoYXNpYQ==

	keyBase32 := base32.StdEncoding.EncodeToString([]byte(secret))

	fmt.Println(keyBase32)

	fmt.Println(key)

	rahasia, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(rahasia))

	// sha1 => one way encryption
	// kita cuma bisa encrypt, tapi tidak bisa decrypt
	// konsep comparing antara plain text dan encrypt text
	// private key sama secret key

	var sha = sha1.New()
	sha.Write([]byte(secret))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)
	// f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8
}
