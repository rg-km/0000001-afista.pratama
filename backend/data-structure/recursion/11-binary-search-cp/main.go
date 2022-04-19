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

/*
		5
	   / \
	  3   8
	 /\   /\
	2  4 7  9
	    /     \
		6	  10

*/

//Recursive Binary Search
func BinarySearch(numList []int64, key int64) int {
	low := 0
	high := len(numList) - 1

	if low <= high {
		// TODO: answer here
	}
	return 0
}
