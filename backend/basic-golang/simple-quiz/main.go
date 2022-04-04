package main

import (
	"log"
	"os"
)

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeText(fileName, data string) {
	file, err := os.Create(fileName)
	errHandler(err)

	defer file.Close()

	_, err = file.WriteString(data)
	errHandler(err)
}

func main() {
	file := "data.txt"
	data := "quiz berhasil dijawab"

	// call func
	writeText(file, data)
}
