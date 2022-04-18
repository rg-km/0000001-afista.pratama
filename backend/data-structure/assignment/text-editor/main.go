package main

import (
	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack // stack = undo
	RedoStack *stack.Stack // stack = redo
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

// "hello"
// 'h' => rune => int32 => angka
// clue : akan lebih enak menggunakan stack

// push
func (te *TextEditor) Write(ch rune) {
	// redo undo
	// redo = nil
	// undo = kita tambah dengan data yang skrg

	te.RedoStack.SetToEmpty()
	te.UndoStack.Push(ch)

	// hell
	//h e l l
	//  o
}

// peek
func (te *TextEditor) Read() []rune {
	// kita ngembaca dari list data uang ada di undo
	var result []rune

	for te.UndoStack.Top > -1 {
		char, _ := te.UndoStack.Pop()

		result = append(result, char)
	}

	// hello
	// o l l e h

	// reverse
	// h e l l o

	// reverse 1
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// reverse 2
	// var reverse []rune

	// for i := len(result); i > 0; i-- {
	// 	reverse = append(reverse, result[i-1])
	// }

	return result
}

// pop
func (te *TextEditor) Undo() { // ctrl + z
	// undo :hello, redo : nil
	// undo: hell, redo : o
	// undo: hel, redo : lo

	char, err := te.UndoStack.Peek()
	if err != nil {
		return
	}

	te.UndoStack.Pop()
	te.RedoStack.Push(char)
}

// push
func (te *TextEditor) Redo() {
	// undo: hel, redo : lo
	// undo: hell , redo : o
	// undo: hello , redo: nil

	// error / stop / not doing
	char, err := te.RedoStack.Peek()
	if err != nil {
		return
	}

	te.RedoStack.Pop()
	te.UndoStack.Push(char)
}
