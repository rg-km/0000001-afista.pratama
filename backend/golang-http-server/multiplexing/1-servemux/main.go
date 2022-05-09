package main

import (
	"fmt"
	"log"
	"net/http"
)

// Pada implementasi handler sebelumnya, ketika kita hit ke endpoint manapun maka handler yang dieksekusi sama.
// Multiplexer berguna untuk melakukan route antara endpoint dengan handler
// Do golang, untuk mengimplementasikan multiplexer kita dapat menggunakan http.ServeMux

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "custom handler mux")
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		// check method allow GET
		if r.Method != "GET" {
			log.Println("Method not allowed")
			fmt.Fprint(w, "Method not allowed")
			return
		}

		next.ServeHTTP(w, r)
	})

}

func main() {
	// Objek http.ServeMux yang bertindak sebagai multiplexer
	mux := http.NewServeMux() // otomatis membuat mux baru

	//mx := http.DefaultServeMux

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home")
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	myHandler := MyHandler{}
	mux.Handle("/handle", myHandler)

	var handler http.Handler = mux

	handler = middleware(handler)

	// set server
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Server run on port 8080")
	//http.ListenAndServe("localhost:8080", mux)
	log.Fatal(server.ListenAndServe())
}
