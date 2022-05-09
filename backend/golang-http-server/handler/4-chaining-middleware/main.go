package main

import (
	"fmt"
	"net/http"
)

// middleware di tengah
// req -> middleware (handler) -> handler
// req -> middleware1 (handler) -> middleware2 (handler) -> middleware3 (handler) -> handler

func secondMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Second middleware: check method GET")
		if r.Method != "GET" {
			http.Error(w, "Method is not allowed.", http.StatusMethodNotAllowed)
			return
		}

		//fmt.Println("Second middleware: check method GET again")
		next.ServeHTTP(w, r)
	})
}

func firstMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("First middleware")

		//fmt.Println("First Middleware Again")
		next.ServeHTTP(w, r)
	})
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world handler")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func main() {
	handler := http.HandlerFunc(helloWorldHandler)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: firstMiddleware(secondMiddleware(handler)),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
