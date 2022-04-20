package main

func squareWorker(input <-chan int, output chan<- int) {
	for i := 1; i <= <-input; i++ {
		//mengirim ke channel output hasil	pangkat 2 nilai dari channel input
		// TODO: answer here
		//output <- <-input * <-input
		output <- (i * i)
	}
}

//mengirim 1-n angka ke squareWorker
func numberGenerator(n int, input chan<- int) {
	for i := 1; i <= n; i++ {
		// TODO: answer here
		input <- i
	}
}
