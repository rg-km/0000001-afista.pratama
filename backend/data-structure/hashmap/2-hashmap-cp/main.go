// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

package main

import "fmt"

func main() {
	var str1 = "fried"
	var str2 = "fired"
	fmt.Println(AnagramsChecker(str1, str2))
}

func AnagramsChecker(str1 string, str2 string) string {
	//return ""

	// 1 mengecek panjangnya dari str1 dan str2 , kalau beda sudah pasti tidak anagram : "Bukan Anagram"

	if len(str1) != len(str2) {
		return "Bukan Anagram"
	}

	// 2 ditata ke dalam sebuah map
	// kita tata huruf2nya
	var map1 = make(map[int32]int)
	var map2 = make(map[int32]int)

	for i := 0; i < len(str1); i++ {
		var key = int32(str1[i])

		if _, ok := map1[key]; ok {
			map1[key]++
		} else {
			map1[key] = 1
		}
	}

	for i := 0; i < len(str2); i++ {
		var key = int32(str2[i])

		if _, ok := map2[key]; ok {
			map2[key]++
		} else {
			map2[key] = 1
		}
	}

	// knee , kenn
	// keen , knee

	// tidak dicek, kalau total tiap2 huruf sama, berarti "Anagram"
	for key, value := range map1 {
		if map2[key] != value {
			return "Bukan Anagram"
		}
	}
	// kalau tidak berarti "Bukan Anagram"

	return "Anagram"
}
