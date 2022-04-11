package main

import "fmt"

func main() {
	var cars1 = []string{"Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"}
	var cars2 = []string{"Ford", "BMW", "Audi", "Mercedes"}

	// ["Ford", "BMW"]

	result, err := SearchMatch(cars1, cars2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Matched car brand: ", result)
	}
}

func SearchMatch(arr1 []string, arr2 []string) ([]string, error) {
	temp := []string{} // ["b"]

	// search / matching

	// ["a", "b", "c"] | ["b", "d", "e"]

	// v = "a"
	// v2 = "b"
	// v2 = "d"
	// v2 = "e"

	// interasi 2
	// v = "b"
	// v2 = "b"
	// v2 = "d"
	// v2 = "e"

	// interasi 3
	// v = "c"
	// v2 = "b"
	// v2 = "d"
	// v2 = "e"

	//O(n) | O(n^2)

	for _, v := range arr1 {
		for _, v2 := range arr2 {
			if v == v2 {
				temp = append(temp, v)
			}
		}
	}

	if len(temp) == 0 {
		return nil, fmt.Errorf("no match")
	}

	return temp, nil
}
