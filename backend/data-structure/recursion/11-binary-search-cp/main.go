package main

import "fmt"

func main() {
	numList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := int64(9)
	fmt.Println(BinarySearch(numList, key))
}

// O(n), O(n^m)

// O(log n)

// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10

// [1, 2, 3, 4, 5] [6, 7, 8, 9, 10] // 1 ops
//					[6,7] [8,9,10] //  1 ops
// 							[8] [9,10] // 1 ops

// 3 ops
// kalau menggunakan looping, 9 ops

/*
		5
	   / \
	  3   8
	 /\   /\
	2  4 7  9
	    /     \
		6	  10

*/

// angka ada, return 1 // true
// nggk ada , return 0 // false

//Recursive Binary Search
func BinarySearch(numList []int64, key int64) int {
	low := 0
	high := len(numList) - 1

	if low <= high {
		// mencari angka 7
		// [3,4,5,6,7,8] , 6 + 0 / 2, idx 3 => 6
		//  [3,4,5] [6,7,8]
		// [6,7,8] => 2 + 0 / 2 , idx 1 => 7
		// [6] [7,8]

		// [7,8] [7] dan [8]

		mid := (high + low) / 2
		if numList[mid] > key {
			return BinarySearch(numList[:mid], key)
		} else if numList[mid] < key {
			return BinarySearch(numList[mid+1:], key)
		} else {
			return 1
		}
	}

	return 0
}
