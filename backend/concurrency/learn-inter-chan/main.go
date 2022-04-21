package main

import (
	"fmt"
	"time"
)

func main() {

	// chan
	// ch := make(chan int)

	// seperti queue

	// tentukan size

	// buffer channel
	// range buffer channel, channel biasa
	// close channel
	// select

	// 3 size [nil, nil, nil]
	// [1,2, nil]

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1
		}
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}

	// close

	// close([1,2, nil]) <- 3 , panic
	// close([1,2, nil]) <- close , panic
	// 1 <- close([2, nil, nil])
	// 2 <- close([nil, nil, nil])

	// select channel

	// main <- beberapa channel, c1, c2, c3, dst
	// c1 := make(chan int) // delay 200 ms
	// c2 := make(chan int) // delay 100 ms + 200
	// c3 := make(chan int) // delay 50 ms + 200

	start := time.Now()

	c1 := make(chan string, 5)
	c2 := make(chan string, 5)

	go func() {
		// tugas 1
		for {
			time.Sleep(300 * time.Millisecond)
			c1 <- "every 300 millisecond"
		}
	}()

	go func() {
		// tugas 2
		for {
			time.Sleep(1000 * time.Millisecond)
			c2 <- "every 1 second"
		}

	}()

	for {
		// fmt.Printf("%s channel 1. At time %s\n", <-c1, time.Since(start)) // delay 300ms
		// fmt.Printf("%s channel 2. At time %s\n", <-c2, time.Since(start)) // delay 1s

		// mana dulu channel yang keisi dari beberapa channel
		select {
		case c := <-c1:
			fmt.Printf("%s channel 1. At time %s\n", c, time.Since(start))
		case c := <-c2:
			fmt.Printf("%s channel 2. At time %s\n", c, time.Since(start))
		default:
			time.Sleep(200 * time.Millisecond)
			fmt.Println("waiting for channel")
		}
	}

	// go 1 -------- go 2
	//  300          delay 1s call
	// delay 1s call
	// 	delay			2s call
	// delay 2s call

	// go 1 -------- go 2
	// 300 call
	// 600 call
	// 900 call
	//				1000 call
	// 1200 call
	// 1500 call
	// 1800 call
	// 				2000 call

	// ch := make(chan int, 1)

	// go func() {
	// 	for i := 1; i <= 5; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(<-ch)
	// }

	// for c := range ch {
	// 	fmt.Println(c)
	// }

	// var arr = []int{1,2,3,4,5}

	// for a := range arr {

	// }

	// main ------------------ go func()
	// set buffer chan (1)
	// <- ch wait
	// 							ch <- 1
	// print 1
	// 							ch <- 2
	// print 2
	// 							ch <- 3
	// print 3
	// ......

	// -----------------------------------

	// main stop

	// ch := make(chan int, 2) // [nil, nil]

	// ch <- 1           //[1, nil]
	// ch <- 2           //[1, 2]
	// fmt.Println(<-ch) //1 <- [2, nil]

	// ch <- 3 // [2, 3] <- 3

	// fmt.Println(<-ch) //2 <- [3, nil]
	// fmt.Println(<-ch) //3 <- [nil, nil]
}

// 5 concurrency -> 1 chan sama
// 1 concurrency , yang mengolah data dari chan

// [5,4,3,2,1]
// [10,9,8,7,6,5,4,3,2] => 1
// [15,14,13,12,11,10,9,8,7,6,5,4,3] => 2

// set size 5 (buffer channel)
// [5,4,3,2,1]
// [6,5,4,3,2] => 1
// [7,6,5,4,3] => 2
// ....
