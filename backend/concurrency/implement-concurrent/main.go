package main

import (
	"encoding/csv"
	"log"
	"os"
)

// multiple request, ke beberapa list website
// concurrent

// http.Get
// tampung csv
// buat output response

// concurrent untuk melakukan multiple request

// crawler

// worker [set max goroutine process, pakai channel]
// dapatkan, terus simpan jadi output response

// worker = limit goroutine process

// 192.0.0.1
// 192.0.0.2
// dst
// 255.255.0.1
// dst

// initialize
func main() {

}

func OpenCSV(fileName string) [][]string {
	// open file
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal("error :", err)
	}

	defer file.Close()

	// reading csv
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("error :", err)
	}

	// return
	return records
}

func StartProcess(records [][]string) {
	// goroutine
}

// crawler
// goroutine
// channel
// set max worker / di limit goroutine
func GetResponseWebsite() {

}
