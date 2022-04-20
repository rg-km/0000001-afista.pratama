package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		//kenapa valuenya sama semua ?
	}

	// sequencial ------ concurrency
	// main
	// loop
	// i = 0 	--- go func(0)
	// i = 1 			 		---go func (1)
	// i = 2									 --- go func(2)
	// ....dst												go ...dst
	// i = 5 		i = 0			i = 1			i = 2 , dst

	time.Sleep(10 * time.Millisecond)
	fmt.Println("main stop")
}
