// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import (
	"fmt"
)

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
	fmt.Println(Intersection(str1, str2))
}

// matematic => himpunan
// datanya unique

func Intersection(str1, str2 []string) (inter []string) {
	//return []string{} // TODO: replace this

	// hash => set
	// str1 => kita tampung ke hash
	// comparing dengan str2
	// kalau ada yang nggk ada di str1,
	// maka tambahkan value ke str1
	set := make(map[string]bool)
	/*
		{
			a = true
			b = true
			c = true
		}
	*/

	// [a,b,c]
	for _, str := range str1 {
		set[str] = true
	}

	// comparing dengan str2
	for _, str := range str2 {
		if set[str] {
			inter = append(inter, str)
		}
	}

	inter = RemoveDuplicates(inter)

	return
}

func RemoveDuplicates(elements []string) (nodups []string) {
	// a,a,b,b,c
	// a,b,c
	set := make(map[string]bool)

	/* set
	{
		a = true
		b = true
		c = true
	}
	*/

	// O(n)
	for _, e := range elements {
		if !set[e] {
			nodups = append(nodups, e)
			set[e] = true
		}
	}

	// O(n*2)
	// for _, e := range elements {
	// 	if _, ok := set[e]; !ok {
	// 		set[e] = true
	// 	}
	// }

	// for key, _ := range set {
	// 	nodups = append(nodups, key)
	// }

	return
}
