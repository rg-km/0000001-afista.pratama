package main

import (
	"fmt"
	"time"
)

// goroutine

//chan<- string => send only channel
// <-chan string => receive only channel

func getDataFromCallAPI(ch chan<- string) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("data dari API terproses")

	ch <- "process 1 done"
}

func getDataFromCallDB(ch chan<- string) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("data dari DB terproses")

	ch <- "process 2 done"
}

// parameter channelnya cuma bisa menerima pesan
// get value
// tapi tidak bisa membaca atau melempar pesan

func getDataFromCallFile(ch chan<- string) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("data dari File terproses")

	ch <- "process 3 done"
}

// 3
// main 		A			B 			C
// channel
// 			   ch<-		  ch<-		   ch<-
// <-ch
// <-ch
// <-ch

// channel buffer
// condition if else, switch case

func main() {
	ch := make(chan string)
	// 1 "process 1 done"
	// 2 "process 2 done"
	// 3 "process 3 done"

	// channel as parameter function
	// cuma bisa menangkap value aja
	// cuma bisa membaca value aja

	startTime := time.Now()

	go getDataFromCallAPI(ch)

	go getDataFromCallDB(ch)

	go getDataFromCallFile(ch)

	fmt.Println("process selesai")

	fmt.Println("execute : ", time.Since(startTime))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	time.Sleep(1 * time.Second)

	// channel

	// channel hanya bisa komunikasi di goroutine
	// kalau tidak, akan fatal error
	// communicating with goroutine
	// blocking execute
	// channel as parameter function

	// CSP
	// message passing passing

	// declare
	// ch := make(chan int)

	// go func() {
	// 	ch <- 10 // menerima message
	// }()

	// go func() {
	// 	ch <- 12 // menerima message
	// }()

	// mendapatkan value
	//val := <-ch

	// membuang value
	//<- ch

	// concurrency, syntax : go
	// runtime.GOMAXPROCS(100)

	// kalau loadnya kecil, 2 core
	// kalau 100 core

	// memproses data
	// menginginkan get data dari API // 100ms
	// get akses ke DB // 100 ms
	// get akses ke file // 100 ms

	// 300ms => 100 ms
	// process

	// fmt.Println("start")

	// go func() {
	// 	fmt.Println("concurrency running")
	// }()

	// go func() {
	// 	fmt.Println("concurrency 2")
	// }()

	// fmt.Println("end")

	// time.Sleep(time.Millisecond * 100) // delay 100 ms
}

// main ----------------    func()1   ---------	func()2
// print start // 1 ms    -+ delay 10ms		-+ delay 10ms
// print end // 1 ms

// setelah delay 10 ms   print concur..     print concur2 ..
// stop
