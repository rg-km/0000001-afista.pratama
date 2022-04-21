package main

import (
	"fmt"
	"time"
)

// goroutine
// rule harus invoke sebuah function

func main() {
	go func() {
		fmt.Println("hello from goroutine")
	}()

	// IIFE
	// immediately invoke function expression

	time.Sleep(10 * time.Millisecond)
	fmt.Println("main stop")
}
