// Sort array terlebih dahulu, kemudian rotasi ke kiri sesuai dengan nilai yang telah ditentukan.
//
// Contoh Sort array:
// Input: [4,5,2,1,3]
// Output: [1,2,3,4,5]

//Contoh RotateLeft:
//Input: 4, [1,2,3,4,5]
//Output: [5,1,2,3,4]

// Explanation RotateLeft:
// untuk melakukan rotasi kiri dengan nilai 4, array mengalami urutan perubahan berikut:
// [1,2,3,4,5] -> [2,3,4,5,1] -> [3,4,5,1,2] -> [4,5,1,2,3] -> [5,1,2,3,4]

// 1.000.000
// overkill
// merge sort, quick sort, counting, radix, dll (complex) => setup (10 detik), running (max 6 detik) , 16 detik
// bubble sort  => setup (1 detik), sorting (3 menit) , 3 menit 1 detik

package main

import "fmt"

func main() {
	var nums = []int{4, 5, 2, 1, 3}
	arrSorted := Sort(nums)
	fmt.Println(arrSorted)
	rotateLeft := RotateLeft(2, arrSorted)
	fmt.Println(rotateLeft)
}

func Sort(arr []int) []int {
	//return []int{}
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}

func RotateLeft(d int, arr []int) []int {
	// [1,2,3,4,5] => [1], [2,3,4,5] => [2,3,4,5] + [1] = [2,3,4,5,1]
	// [2,3,4,5,1] => [3,4,5,1] + [2] = [3,4,5,1,2]
	// [3,4,5,1,2] => [4,5,1,2] + [3] = [4,5,1,2,3]

	for i := 0; i < d; i++ {
		arr = append(arr[1:], arr[0])
	}
	return arr
}
