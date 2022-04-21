package main

import "fmt"

//buferred channel tidak blocking jika buffer belum penuh
func main() {
	// cap(), size, space
	input := make(chan string, 3) // buffer channel [nil, nil, nil]

	input <- "halo"          // ["halo", nil, nil]
	input <- "dari buffered" // ["halo",  "dari buffered", nil]
	input <- "channel"       // ["halo",  "dari buffered", "channel"]

	//<-input // queue : "halo" <- ["dari buffered", "channel", nil]

	//input <- "extend space" // error fatal / panic : [ "dari buffered", "channel",  "extend space"]

	// len, banyaknya elemen

	//fmt.Println(len(input)) // 2

	// len(input) // 1

	for i := 0; i < len(input); i++ {
		fmt.Println(<-input)
	}
}
