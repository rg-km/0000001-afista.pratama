package main

import (
	"fmt"
)

//mengembalikan hasil pangkat dua angka 1-10
//dapat dilihat tidak terjadi blocking selama buffer belum penuh
func squareWorker(output chan<- int) {
	for i := 1; i < 11; i++ { // 1- 10
		output <- i * i
	}
	fmt.Println("selesai mengirim")
}

func main() {
	output := make(chan int, 10)
	// [nil, nil, nil ... 10]
	// [1, 4, 9, 16, ..... 100]

	go squareWorker(output)
	//time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10; i++ {
		fmt.Println("main receive output:", <-output)
	}
}
