package main

import (
	"fmt"
	"time"
)

//goroutine dapat berjalan secara concurrent
//disini goroutine akan mengcapture variable i secara reference
//ini perlu dilakukan agar goroutine dapat mengubah value i
//bisa disebut closure juga karena fungsi goroutine reference variable di luar scopenya
func main() {
	i := 0 //nilai i akan diubah oleh goroutine
	go func() {
		for i < 10 {
			fmt.Println("increment", i)
			i++
		}
	}()

	for i < 10 {
		time.Sleep(10 * time.Microsecond)
		fmt.Println("loop")
	}

	fmt.Println("main stop")

	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }
}

// main  		------- go loop
// start loop
// loop
// loop 		10ms increment 1
// loop			increment 2
// loop 		increment 3
// .....		....
// main stop

// function generator

// func
// execute 1
// execute 2
// execute 3
// pause

// closure
// encapsulate data in function
func adder() func(int) int {
	// encapsulate
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// func incrementData() func() int {
// 	var i int = 0

// 	return func() int {
// 		i++
// 		return i
// 	}
// }
