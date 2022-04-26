package main

import (
	"fmt"
	"sync"
	"time"
)

// type count struct {
// 	mu *sync.Mutex
// }

// func (c *count) lock() {
// 	//fmt.Printf("mutex address: %p\n", &c.mu)
// 	c.mu.Lock()
// }

// func (c *count) unlock() {
// 	//fmt.Printf("mutex address: %p\n", &c.mu)
// 	c.mu.Unlock()
// }

func main() {
	// channel, untuk komunikasi menggunakan share memori ke share variable
	// mutex,

	//mu := &sync.Mutex{} // unlock Mutex
	// var a int

	// c := count{mu: &sync.Mutex{}}

	// ///mu.Unlock()

	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		c.lock()
	// 		a++
	// 		c.unlock()
	// 	}()
	// }

	// go1 --- go2 --- go3 --- go4 --- go5 .. dst
	// lock    ------blocking-------------
	// a++
	// unlock  lock --------- blocking ------
	//			a++
	// 			unlock   lock -----blocking ----
	// 					  a++
	// 					unlock ....dst

	// time.Sleep(400 * time.Millisecond)

	// fmt.Println("total", a)

	// deadlock

	// c1 := make(chan int, 1)
	// c2 := make(chan int, 1)

	// go func() {
	// 	fmt.Println("goroutine 1")
	// 	c1 <- 1
	// 	val := <-c2
	// 	fmt.Println("val go 1", val)

	// }()

	// go func() {
	// 	fmt.Println("goroutine 2")
	// 	c2 <- 1
	// 	val := <-c1

	// 	fmt.Println("val go 2", val)
	// }()

	// time.Sleep(100 * time.Millisecond)

	// starvation
	done := make(chan bool, 1)
	var mu sync.Mutex
	counterGreedy := 0
	counterPolite := 0

	// goroutine 1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				counterGreedy++
				time.Sleep(100 * time.Microsecond)
				//do something to shared resource
				mu.Unlock()
			}
		}
	}()

	// goroutine 2
	for i := 0; i < 1000; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		counterPolite++
		//do something to shared resource
		mu.Unlock()
	}
	done <- true
	fmt.Println("counter greedy:", counterGreedy)
	fmt.Println("counter polite:", counterPolite)

}
