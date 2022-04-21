package main

import (
	"fmt"
	"time"
)

func main() {
	output := make(chan string, 10)

	go sender(output)
	//time.Sleep(100 * time.Millisecond)
	for msg := range output {
		fmt.Println(msg)
	}
}

func sender(output chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("send to channel", i)
		output <- "every 200 millisecond"
		time.Sleep(200 * time.Millisecond)
	}

	close(output)
	//fmt.Println("channel close")
}

//jika menggunakan for ... range channel
//kita perlu menutup channel tersebut saat sudah tidak ada data yang dikirim
//yang menutup channel sebaiknya adalah pengirim data (disini fungsi sender)
