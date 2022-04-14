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
	UndoStack *stack.Stack
	RedoStack *stack.Stack
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

func (te *TextEditor) Write(ch rune) {
	// TODO: answer here
}

func (te *TextEditor) Read() []rune {
	// TODO: answer here
}

func (te *TextEditor) Undo() { // ctrl + z
	// hello
	// hell
	// hel
	// TODO: answer here
}

func (te *TextEditor) Redo() {
	// hel
	// hell
	// hello

	// error / stop / not doing

	// TODO: answer here
}
