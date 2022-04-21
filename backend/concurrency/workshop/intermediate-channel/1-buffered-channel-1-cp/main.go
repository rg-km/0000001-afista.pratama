package main

//buffered channel tidak blocking selama buffer belum full
func testBuffer(output chan int) {
	input := make(chan int, 4)
	input <- 0
	input <- 1
	input <- 2
	input <- 3

	sum := 0
	for i := 0; i < cap(input); i++ {
		// TODO: answer here
		sum += <-input // [0, 1, 2, 3]
		// loop1
		// sum + 0  : input [1,2,3, nil]
		// sum + 1  : input [2,3, nil, nil]
		// dst ....
	}

	output <- sum
}
