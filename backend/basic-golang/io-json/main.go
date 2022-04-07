package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func handleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type Student struct {
	Name  string `json:"nama"` //annotation
	Score int    `json:"nilai"`
	Class string `json:"kelas"`
}

func writeJSON(data []Student) {
	pathFile := "data.json"

	// jsonData, err := json.Marshal(students) // marshal , ngubah dari struct / object => JSON

	jsonDataIndent, err := json.MarshalIndent(data, "\t", "")
	handleErr(err)

	err = ioutil.WriteFile(pathFile, jsonDataIndent, 0644)
	handleErr(err)

	fmt.Println("berhasil membuat file json")
}

func readJSON() []Student {
	pathFile := "data.json"

	file, err := os.OpenFile(pathFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	handleErr(err)

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	handleErr(err)

	var students []Student

	err = json.Unmarshal(data, &students)
	handleErr(err)

	return students
}

// ingin mengubah data
// readJSON() => baca dulu file json yang ada
// diubah2 => nambah 1 file
// writeJSON() write data ke json yang ada

func main() {
	students := readJSON()

	// proses update
	students = append(students, Student{
		Name:  "si D",
		Score: 100,
		Class: "BE02",
	})

	writeJSON(students)
}
