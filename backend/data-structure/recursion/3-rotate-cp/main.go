// Expected Rosated:
// [1 2 3 4 5 6 7 8 9 10]
// [(1) 2 3 4 5 6 7 8 9 (10)] => [(10) 2 3 4 5 6 7 8 9 (1)]
// [10 2 3 4 5 6 7 8 9 1]
// [10 (2) 3 4 5 6 7 8 (9) 1] => [10 (9) 3 4 5 6 7 8 (2) 1]

// [10 2 3 4 5 6 7 8 9 1]
// [10 2 (3) 4 5 6 7 (8) 9 1]
// [10 2 3 (4) 5 6 (7) 8 9 1]

// Proses Rotate:
// [1 2 3 4 5 6 7 8 9 10]=> 10 dan 1
// [10 9 3 4 5 6 7 8 2 1] => 9 dan 2
// [10 9 8 4 5 6 7 3 2 1] => 8 dan 3
// [10 9 8 7 5 6 4 3 2 1] => 7 dan 4
// [10 9 8 7 6 5 4 3 2 1] => 6 dan 5
// [10 9 8 7 5 6 4 3 2 1] => 5 dan 6
// [10 9 8 4 5 6 7 3 2 1] => 4 dan 7
// [10 9 3 4 5 6 7 8 2 1] => 3 dan 8
// [10 2 3 4 5 6 7 8 9 1] => 2 dan 9

package main

import "fmt"

func Rotate(index int, args []int) []int {
	if len(args) == 0 {
		return []int{}
	}

	if index == -1 {
		return args
	}

	if index == 0 {
		// TODO: answer here
		// swap
		args[0], args[len(args)-1] = args[len(args)-1], args[0]
		return Rotate(1, args)
	}

	if index == len(args)-1 {
		return Rotate(-1, args)
	}

	return Rotate(index+1, args)

}

func main() {
	mylist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rotated := Rotate(0, mylist)
	fmt.Println(rotated)
}
