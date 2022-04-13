// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired" //list huruf sama , anagram
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
	// return ""
	// "apple" 5 | "paddle" 6
	// 1 mengecek panjangnya dari str1 dan str2 , kalau beda sudah pasti tidak anagram : "Bukan Anagram"

	if len(str1) != len(str2) {
		return "Bukan Anagram"
	}

	// 2 ditata ke dalam sebuah map
	// kita tata huruf2nya

	// str1 = "keen" , str2 = "knee"

	/*
		str1 {
			"k" : 1,
			"e" : 2,
			"n" : 1
		}

		str2 {
			"k" : 1,
			"n" : 1,
			"e" : 2
		}
	*/

	var map1 = make(map[int32]int)
	// var map1 = make(map[string]int)
	var map2 = make(map[int32]int)

	for _, val := range str1 {
		map1[val]++
		// map1[string(val)]++
	}

	for _, val := range str2 {
		map2[val]++
	}

	// knee {k:1, n:1, e:2}, kenn {k:1, e:1, n:2}
	// keen , knee

	// tidak dicek, kalau total tiap2 huruf sama, berarti "Anagram"
	for key, value := range map1 {
		// {"k" : 1,"e" : 2,"n" : 1}

		if map2[key] != value {
			return "Bukan Anagram"
		}
	}
	// kalau tidak berarti "Bukan Anagram"

	return "Anagram"
}
