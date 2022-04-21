package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	//sumAll([]int{0, 1, 2, 3}) // 6
	//sumAll([]int{1, 2, 3, 0}) // 6

	// worst case testing, defensive programming
	// fmt.Println(findMin([]int{})) //nil
	// fmt.Println(findMin([]int{100_000_000_000_000, 1, 2, 3}))
	// fmt.Println(findMin([]int{-100_000_000_000_000, 1, 2, 3}))

	// golang bahasa strict
	// tipe datanya harus dideclare di awal
	// -2_147_483_648 through 2_147_483_647

	// 2_147_483_648 = -2_147_483_648
	//var arr = []int{32, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 7, 3, 9, 11, 15, 17, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

	fmt.Println("total price", shoppingPrice())
}

// clean architecture
// model <- repository -> service -> I/O : controller, layer output

// debugging itu apa
// handle debugging
// - memahami error (semantic & syntax error)
// - error : compile error & runtime error
// - strategic : SRP -> memecah tiap fungsi menjadi hanya 1 tugas
// tips :
// - simplify test (data input, codingan)
// - divide & conquer (binary tree) - pecah2
// - menggunakan logging
// - selesaikan bug 1 persatu, jangan sampai memunculkan bug baru setelah solving bug sebelumnya.
// - bertanya, dan istirahat

// case shoppingPrice
func shoppingPrice() int {
	totalItem, err := getTotalItem()
	if err != nil {
		log.Println("error :", err.Error())
	} else {
		log.Println("function getTotalItem, success")
	}

	prices := getPrices(totalItem)
	return getTotalPrice(prices)
}

func getTotalItem() (int, error) {
	var totalItem int
	fmt.Print("Total item: ")
	fmt.Scanf("%d", &totalItem)

	if totalItem == 0 {
		return 0, errors.New("input invalid")
	}

	return totalItem, nil
}

func getPrices(totalItem int) []int {
	prices := make([]int, totalItem)
	for i := 0; i < totalItem; i++ {
		fmt.Printf("Harga %d :", i+1)
		fmt.Scanf("%d", &prices[i])
	}

	return prices
}

func getTotalPrice(prices []int) int {
	var totalPrice int

	for i := 0; i < len(prices); i++ {
		totalPrice += prices[i]
	}

	return totalPrice
}

func twoSum(arr []int, target int) (int, int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if arr[i]+arr[j] == target && i != j {
				return arr[i], arr[j]
			}
		}
	}
	return -1, -1
}

// func countDiscount(total float64) float64 {
// 	var discountRate float64 = 0.5
// 	return total * discountRate
// }

// func sum(num1, num2 int) int {
// 	//return num1 * num2
// 	return num1 + num2
// }

// func oddNumber(limit int) {
// 	num := 1
// 	while (num < limit) {
// 		fmt.Println(num)
// 		num += 2
// 	}
// }

// func hi() {
// 	greeting := "welcome to"
// 	nama := "ruang" // declare + arrign var nama = "ruang"
// 	nama := "guru" // compile error, syntax error
// 	fmt.Println(greeting, nama)
// }

func findMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

func sumAll(arr []int) int {
	fmt.Println("Processing: ", arr)

	sum := 0
	// fix initial loop, / i , from 1 to 0 (same as index)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println("Result: ", sum)
	return sum
}
