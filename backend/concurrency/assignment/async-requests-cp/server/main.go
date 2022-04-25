package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ruang-guru/playground/backend/concurrency/assignment/async-requests-cp/server/handler"
)

func main() {
	http.HandleFunc("/", handler.GetMessage)

	fmt.Println("starting server")
	fmt.Println("running on server: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
