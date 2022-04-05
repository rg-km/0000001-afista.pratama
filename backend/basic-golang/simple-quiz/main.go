package main

import (
	"log"
	"os"
)

// handling an error
func errHandler(err error) {
	if err != nil {
		log.Fatal(err) // throw some error with log
	}
}

func writeText(fileName, data string) {
	// create a file from param value
	file, err := os.Create(fileName)
	errHandler(err)

	defer file.Close() // close connection

	_, err = file.WriteString(data)
	errHandler(err)
}

func main() {
	file := "data.txt"
	data := "quiz berhasil dijawab"

	// call func
	writeText(file, data)
}
