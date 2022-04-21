package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	Size int
	Top  int
	Data []int
}

func NewStack(size int) Stack {
	return Stack{
		Size: size,
		Top:  -1,
		Data: []int{},
	}
}

func (s *Stack) Push(Elemen int) error {
	if s.Size == len(s.Data) {
		return ErrStackOverflow
	}

	s.Top++
	s.Data = append(s.Data, Elemen)

	return nil
}
