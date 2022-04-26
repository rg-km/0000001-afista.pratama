package main

import (
	"log"
	"sync"
)

func testWG(output chan<- []bool) {
	// TODO: answer here
	var wg sync.WaitGroup

	done := make([]bool, 5)

	wg.Add(5)

	log.Println(wg)

	for i := 0; i < 5; i++ {
		// TODO: answer here
		//wg.Add(1) // 5
		go func(i int) {
			// TODO: answer here
			defer func() {
				wg.Done()
				log.Println(wg)
			}() // -1
			done[i] = true
		}(i)
	}

	// 5
	// go1  --- go2 --- go3 --- go4 --- go5

	// -1 		-1		-1 		-1		-1
	// finish goroutine

	// counter 0

	// TODO: answer here
	wg.Wait() // counter 0
	output <- done
}
