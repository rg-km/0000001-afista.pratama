package main

import (
	"fmt"
)

//saat menggunakan channel sebagai parameter function
//kita bisa menentukan apakah kita ingin mengirim atau menerima
func main() {
	c := make(chan string)

	block := make(chan bool)

	go receive(c)
	go send(c, block)
	//time.Sleep(100 * time.Millisecond)

	block <- true
}

// <-chan disini berarti channel c hanya bisa menerima data
// receive only channel parameter
func receive(c <-chan string) {
	fmt.Println("receive data ")
	fmt.Println(<-c)
	//c <- "hello"
}

// chan<- disini berarti channel c hanya bisa mengirim data
// send only channel parameter
func send(c chan<- string, block chan bool) {
	fmt.Println("send data")
	c <- "hello"
	//<-c
	<-block

}
