package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0

	for i := 0; i < 5; i++ {
		go func(i int) {
			result += (i * i)
			fmt.Println("result", i*i)
		}(i)
	}

	time.Sleep(10 * time.Millisecond)

	fmt.Println(result)
}
