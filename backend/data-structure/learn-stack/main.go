package main

import (
	"errors"
	"fmt"
)

// Assignment
// parentheses-validation
// text-editor

// deadline malam ini jam 23.59

// stack
// struct + array / slice
// - kalau menambahkan data, itu harus dari belakang
// - kalau mau mengambil data, harus diambil dari belakang
// - kita hanya bisa mengakses data di posisi belakang

// top = representasi posisi paling terakhir / belakang di stack
// top = mengikuti index di posisi terakhir array

// [1,2] + 3 => [1,2,3] => ambil => [1,2] , 3

// encapsulation
type stack struct {
	Top  int   // posisi index terakhir // default value 0, "", false
	Data []int // list data di penampung stack
}

type stackWithArray struct {
	Top  int
	Data []int
}

// [nil, nil, nil, nil, nil]
// [1, nil, nil, nil, nil]
// [1, 2, nil, nil, nil]
// [1,2,3,4,5] + 6 => stack overflow

func NewStackWithArray(sizeArray int) *stackWithArray {
	return &stackWithArray{
		Top:  -1,
		Data: make([]int, sizeArray),
	}
}

func NewStack() *stack {
	return &stack{
		Top: -1,
	}
}

func SetStack(data ...int) *stack { // advance concept
	return &stack{
		Top:  len(data) - 1,
		Data: data,
	}
}

// 3 method syarat (wajib)
// push (menambah data) // datanya maksimal, stack overflow
// pop (mengambil data) // datanya kosong, stack undeflow
// peek (melihat data) // data kosong, empty stack

func (s *stack) Push(data int) error {
	// kalau nambah data, maka top + 1, data ditambah ke Data
	if s.Top == len(s.Data) {
		return errors.New("stack overflow")
	} else {
		s.Top += 1
		s.Data = append(s.Data, data)
	}

	return nil
}

func (s *stack) isEmpty() bool {
	return s.Top == -1
}

func (s *stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack underflow")
	} else {
		// [1,2,3,4], top = 3
		// pop : 4 out, [1,2,3], top = 2
		// pop : 3 out, [1,2], top = 1
		// pop : 2 out, [1], top = 0
		// pop : 1 out, [], top = -1
		// pop: stack underflow

		// setiap pop (pola)
		// dapatkan value di index terakhir
		// top berkurang 1
		// array, dipotong 1 data di index terakhir

		// array slice [:len(array) - 1]
		popperData := s.Data[s.Top]

		s.Top -= 1 // s.Top = s.Top - 1 (decrement)
		s.Data = s.Data[:len(s.Data)-1]

		return popperData, nil
	}
}

func (s *stack) Peek() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("data stack empty")
	} else {
		return s.Data[s.Top], nil
	}
}

func main() {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack)

	val, _ := stack.Pop()

	fmt.Println(val, stack)

	peek, _ := stack.Peek()

	fmt.Println(peek, stack)
}
