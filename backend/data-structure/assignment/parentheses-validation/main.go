package main

import (
	"github.com/ruang-guru/playground/backend/data-structure/assignment/parentheses-validation/stack"
)

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

// contoh :4
// input: s = "([)]" => "([])"
// Output : false

// contoh : 5
// input: s = "[ [{}] () [{}] ]"

func IsValidParentheses(s string) bool {
	// var stacks stack.Stack = stack.Stack{
	// 	Top:  -1,
	// 	Data: nil,
	// }

	stacks := stack.Stack{
		Top:  -1,
		Data: nil,
	}

	// ini mudah disolve menggunakan stack
	// package stack

	//"([)" => ""
	if len(s)%2 == 1 || s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		var current rune = rune(s[i])
		if current == '(' || current == '{' || current == '[' {
			stacks.Push(current)
		}
		if current == ')' || current == '}' || current == ']' {
			if stacks.IsEmpty() {
				return false
			}

			inData, _ := stacks.Peek()
			if inData == '(' && current == ')' || inData == '{' && current == '}' || inData == '[' && current == ']' {
				_, _ = stacks.Pop()
			} else {
				return false
			}

		}
	}

	return stacks.IsEmpty()
}

// sek mas e sekk
//wiss loss
