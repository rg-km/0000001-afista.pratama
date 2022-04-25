package main

import (
	"fmt"
	"time"
)

func raceCondition() {
	var a int32

	var i int32

	var input = make(chan int32)

	for i = 1; i <= 1000; i++ {
		go func() {
			//atomic.AddInt32(&a, i)
			input <- 1
		}()
	}

	for i := 0; i < 1000; i++ {
		a += <-input
	}

	// go1 -- go2 -- go3 -- go4 -- .... -- go1000
	// 0	  1
	// 1      3
	// 				3
	//				6
	// 						6
	//						10
	// 								10
	//								15 ... dst

	time.Sleep(1 * time.Second)
	fmt.Println(a)
}

func main() {

	raceCondition()

	// waitgroup (blocking)
	// main dengan counter
	//var wg sync.WaitGroup

	//wg2 := sync.WaitGroup{}

	// Add (menambah counter)
	// Done (decerement, ngurangi 1 counter)
	// Wait (akan menunggu sampai si counter 0)

	// goroutine 1
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()

	// 		fmt.Println("data", i+1)
	// 	}(i)
	// }

	// wg.Wait()

	// fmt.Println("process done")

	// // goroutine 2
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()

	// 		fmt.Println("data go 2", i+1)
	// 	}(i)
	// }

	// wg.Wait()

	// fmt.Println("process 2 done")

	// wg.Add(2) // set counternya 2

	// fmt.Println("wg 1")
	// wg.Done() // counter -1 = 1

	// fmt.Println("wg 2")
	// //wg.Done() // counter -1 = 0

	// wg.Wait() // done // 1

	// fmt.Println("waiting done")

}

func LearnWorker() {
	var imageQueue = make(chan string)
	var videoQueue = make(chan string)
	c := make(chan bool)

	// goroutine 5
	// go func() {
	// 	time.Sleep(200 * time.Millisecond)
	// 	imageQueue <- "image1.jpg"
	// }()
	// go func() {
	// 	time.Sleep(200 * time.Millisecond)
	// 	imageQueue <- "image2.jpg"
	// }()
	// go func() {
	// 	time.Sleep(200 * time.Millisecond)
	// 	imageQueue <- "image3.jpg"
	// }()
	// go func() {
	// 	time.Sleep(400 * time.Millisecond)
	// 	imageQueue <- "image4.jpg"
	// }()
	// go func() {
	// 	time.Sleep(600 * time.Millisecond)
	// 	imageQueue <- "image5.jpg"

	// 	close(imageQueue)
	// }()

	// go receive si channel / thread pool / worker thread
	go func() {
		// for img := range imageQueue {
		// 	fmt.Println(img + "resize")
		// }
		// c <- true

		for {
			select {
			case img := <-imageQueue:
				fmt.Println(img + " resize")
			case vid := <-videoQueue:
				fmt.Println(vid + " resize")
			case <-c:
				break
			}
		}

	}()

	imageQueue <- "image1.jpg"
	imageQueue <- "image2.jpg"
	imageQueue <- "image3.jpg"
	imageQueue <- "image4.jpg"
	imageQueue <- "image5.jpg"
	videoQueue <- "video1.jpg"
	videoQueue <- "video2.jpg"
	videoQueue <- "video3.jpg"
	videoQueue <- "video4.jpg"
	videoQueue <- "video5.jpg"

	// close(imageQueue)
	// close(videoQueue)

	c <- true

	//time.Sleep(1 * time.Second)
}
