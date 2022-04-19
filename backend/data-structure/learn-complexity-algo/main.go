package main

import "fmt"

// complexity algo / analisis

// 2 hal dalam eksekusi kode
// - time complexity / operation
// - space complexity / memory

// big O notation
// O(1), O(n), O(n log n), O(n^m), O(n2), O(log n)

func main() {
	// mencari angka dari sebuah list data
	var list = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(checkNumberInList(list, 15))
}

func checkNumberInList(arr []int, findNumber int) bool {
	//var result bool

	// O(n)
	// for i := 0; i < len(arr); i++ {
	// 	if arr[i] == findNumber {
	// 		result = true
	// 		break
	// 	}
	// }

	// optimize = hashMap O(1), O(log n)
	// buat penampung dari array menjadi hashMap, kemudian kita bisa finding

	var hashMap = map[int]bool{}

	for i := 0; i < len(arr); i++ {
		val := arr[i]

		if _, ok := hashMap[val]; !ok {
			hashMap[val] = true
		}
	}

	return !!hashMap[findNumber]

	//return result
}
