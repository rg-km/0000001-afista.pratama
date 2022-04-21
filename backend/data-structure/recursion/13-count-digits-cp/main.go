// Golang program to count digits of given number
// using recursion

package main

import "fmt"

// var count int = 0

//function to count digits
func CountDigits(num int) int {
	if num > 0 {
		return 1 + CountDigits(num/10) // TODO: replace this
	} else {
		return 0
	}
	//return count
}

// 1441
// 1 + CountDigits(144) 144
// 1 + 1 + CountDigits(14)
// 1 + 1 + 1 + CountDigits(1)
// 1 + 1 + 1 + 1 + CountDigits(0) => 0
// 1 + 1 + 1 + 1 + 0
// 4

// 1441 => 4
// 200_000 = 6

func main() {
	var num int = 0
	var result int = 0

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	result = CountDigits(num)
	fmt.Printf("Count of digits is: %d\n", result)
}
