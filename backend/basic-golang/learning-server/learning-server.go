package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const htmlTemplate = `
<div>
  <h1>  Welcome {{ .Firstname }} </h1>
  <h1>  Your age {{ .Age }} </h1>
</div>
`

type User struct {
	Firstname string
	Age       int
}

func HandleRootRoute(w http.ResponseWriter, r *http.Request) {
	// method apa yang dipakai oleh user
	// request data , method post
	fmt.Println("GET / with template")

	testTemplate, err := template.New("htmlTemplate").Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html")

	err = testTemplate.Execute(w, User{Firstname: "Afis", Age: 23})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleJson(w http.ResponseWriter, h *http.Request) {
	// json
	w.Header().Set("Content-Type", "application/json")

	dataUserJson, err := json.Marshal(map[string]interface{}{
		"name":    "afista",
		"age":     23,
		"address": "jember",
		"tinggi":  172.23,
		"hobbies": []string{"coding", "sleeping", "eating"},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(dataUserJson)
	return
}

func main() {
	// http.ListerAndServe
	handle := http.HandlerFunc(handleJson)

	// middleware
	// req => middleware => handleFunc
	// auth

	// root route , akar router
	http.Handle("/json-data", handle)
	http.HandleFunc("/", HandleRootRoute)

	// http.HandlerFunc // penamaan type ,
	// http.Handle()

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
	// port :8080, 127.0.0.1 / localhost:8080
}
