package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/ruang-guru/playground/backend/concurrency/implement-concurrent/process"
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

var Worker = 200 // 8 // 6 dst

// goroutine max 100
// 100 <- 1 1 1 1 1
// 1 <- 100 <- 1 1 1 1
// 1 1 <- 100 <- 1 1 1

func main() {
	// set gomaxprocs
	runtime.GOMAXPROCS(2)
	start := time.Now()

	datas := OpenCSV("top-100.csv")

	_ = os.Mkdir("response-web", 0755)

	process.AsyncStartProcess(Worker, datas, "top-100-resp.csv")

	fmt.Println("process :", time.Since(start))
}

func OpenCSV(fileName string) [][]string {
	// open file
	log.Println("trying open file", fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal("error :", err)
	}

	defer file.Close()

	// reading csv
	log.Println("reading file", fileName)
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("error :", err)
	}

	// return
	log.Println("success get file", fileName)
	return records
}
