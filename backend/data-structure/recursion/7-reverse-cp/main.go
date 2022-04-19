// Reverse

package main

import "fmt"

// recursive
func Reverse(st []string, depth int) string {
	if depth == -1 { // biar tidak infinite loop  , max 0
		return ""
	}
	str := st[depth] + Reverse(st, depth-1)
	// st[8] = "I" + Reverse(st, 7)
	// "I" + "N" + Reverse(st, 6)
	// "I" + "N" + "D" + Reverse(st, 5)
	// .....
	// Reverse(st, 0) => "I" + "N" + "D" + "O" .... + "A"
	// Reverse(st, -1) => ""
	// "INDONESIA" + "" => "INDONESIA"

	return str
}

// "A", "I", "S", "E", "N", "O", "D", "N", "I",  // depth = 8

func main() {
	str := []string{"A", "I", "S", "E", "N", "O", "D", "N", "I"}
	s := Reverse(str, len(str)-1)
	fmt.Println(s)

	// var rev string

	// for i := len(str) - 1; i >= 0; i-- {
	// 	rev += str[i]
	// }
}
