package simpleserver

import (
	"encoding/json"
	"net/http"
)

// Perlu kalian ketahui bahwa Interface atau signature yang akan kita gunakan disini adalah
// http.HandlerFunc https://pkg.go.dev/net/http#HandlerFunc dan http.Handler https://pkg.go.dev/net/http#Handler

// Decorator pattern di Go biasanya sering digunakan pada web server. Misalnya keperluan logging
type Person struct {
	Name  string `json:"name"` // tag
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// json.Marshal, json.NewEncoder

func New() *Server {
	return &Server{}
}

type Server struct {
}

/*
{
	"name": "John Doe",
	"age": 30,
	"email": "john_doe@gmail.com"
}
*/

func (s *Server) GetPerson(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john_doe@gmail.com",
	}

	// output JSON
	w.Header().Set("Content-Type", "application/json")
	// response code
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

type Logging struct {
}

// Karena agak ribet untuk melakukan testing pada stdout. Maka disini kita menggantinya dengan Header
func (l Logging) AddLogging(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set header dari func handler terus di logging
		w.Header().Set("System-Log", "logged")
		endpoint(w, r)
	})
}
