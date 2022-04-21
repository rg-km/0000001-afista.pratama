package main

func squareWorker(input <-chan int, output chan<- int) {
	//lakukan for range loop
	// TODO: answer here
	for val := range input {
		output <- (val * val)
	}
}

func receiver(output chan<- int) {
	input := make(chan int, 10) // mengirim 0-10 ke squareWorker
	go squareWorker(input, output)
	for i := 0; i < 10; i++ {
		//kirim nilai i ke channel
		// TODO: answer here
		input <- i
		// [0, 1, 4, 9, 16, dst ....]
	}
}
