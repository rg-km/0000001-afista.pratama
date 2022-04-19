// Menghitung Faktorial Menggunakan Recursion dalam Go

package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i

	}
	// 5
	// 5 * 4 = 20 * 3 = 60 * 2 = 120 * 1 = 120

	return result
}

// 5
// 5 * factorialRecursive(4)
// 5 * 4 * factorialRecursive(3)
// 5 * 4 * 3 * factorialRecursive(2)
// 5 * 4 * 3 * 2 * factorialRecursive(1) = 1
// 5 * 4 * 3 * 2 * 1 = 120
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func countingNumberToZero(value int) int {
	// condition
	if value == 0 {
		return 0
	} else {
		return value + countingNumberToZero(value-1)
	}
}

func main() {
	// loop := factorialLoop(10)
	// fmt.Println(loop)

	// recursive := factorialRecursive(10)
	// fmt.Println(recursive)

	res := countingNumberToZero(5) // => 5 + 4+ 3+2+1+0 = 15

	fmt.Println(res)
}
