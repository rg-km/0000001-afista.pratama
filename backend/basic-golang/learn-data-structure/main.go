package main

import "fmt"

func main() {
	// array
	// indexing, append, copy, make, element
	// data structure

	//var numb1 = []int{1, 3, 2, 4, 5}
	// var numb2 = []int{2, 3, 4, 5, 1}
	// var numb3 = []int{100, 233, 24, 15, 2321, 123, 213, 54, 46, 57, 68}

	// fmt.Println(Sorting(numb1))
	// fmt.Println(Sorting(numb2))
	// fmt.Println(Sorting(numb3))

	// sort.Ints(numb3)

	// fmt.Println(numb3)

	// biasanya untuk sorting

	// sort bawaan golang - quick sort

	// map
	// make, key, for range, dll
	// data structure
	// dictionary => O log n => O(1)
	// mendekati O (2^n) , merah

	// unique number [1,1,1,1,2,2,2,2,2,2,3,3,3,3,4,4,4,4,5,5,5,6]
	// [1,2,3,4,5,6]
	test := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 6}
	fmt.Println(GetUniqueNumber(test))

	// array multidimensi [[1,2],[3,4], [5,6]]
}

func GetUniqueNumber(arr []int) []int {
	// hash map => hashmap
	var result []int // 1, 2, 3, 4, 5, 6

	var dict = make(map[int]bool)
	// 1 : true
	// 2 : true

	for _, ar := range arr {
		if _, ok := dict[ar]; !ok {
			result = append(result, ar)
			dict[ar] = true
		}
	}

	return result
}

func Sorting(arr []int) []int {
	// bubble sort

	// start
	// flag = false

	// [2,3,4,5,1]
	// [2,3,4,1,5], flag = true
	// [2,3,1,4,5], flag = true
	// [2,1,3,4,5], flag = true
	// [1,2,3,4,5], flag = true
	// [1,2,3,4,5], flag = false
	// stop logic
	var swap = true

	// Big o => O (n^n)

	for swap {
		swap = false

		for i := 1; i < len(arr); i++ {
			// comparing
			if arr[i-1] > arr[i] {
				// swich
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
	}

	return arr
}
