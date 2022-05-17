package main

import "fmt"

// implementasi / merepresentasikan sebuah tree code
/*
			code1
			/  \
         code2  code3
		 /
		code4

		dst .....
*/

// folder dan file

// 1 common interface , penanda: struct apa saja yang bisa masuk
// method signature

// duck typing
// filosofinya : selama bisa bersuara "cuack" seperti bebek, maka dia bebek, terlepas apapun hewannya

// selama struct memiliki method yang dibutuhkan, maka dia jadi bagian si interface

type Component interface {
	Search(string) bool
}

// 2 struct
type File struct {
	Name      string // ex: data
	Extention string // ex: .txt, .go, .docs
	Data      string // representasi isi datanya di file
}

func (f *File) Search(name string) bool {
	fmt.Println("search file :", f.Name)

	if f.Name == name {
		fmt.Println("Found file:", f.Name)
		return true
	}

	return false
}

type Folder struct {
	Component  []Component // bisa diisi dengan struct File dan Folder
	NameFolder string
}

func (f *Folder) Add(component Component) {
	f.Component = append(f.Component, component)
}

func (f *Folder) Search(name string) bool {
	fmt.Println("search file from folder : ", f.NameFolder)

	for _, component := range f.Component {
		if component.Search(name) {
			return true
		}
	}

	return false
}

// client
func main() {
	file1 := File{
		Name:      "data",
		Extention: ".txt",
		Data:      "Hello World",
	}

	file2 := File{"test", ".txt", "hello hello test"}

	folder1 := Folder{NameFolder: "Folder1"}

	subFolder1 := Folder{NameFolder: "SubFolder1"}

	SubOfSubFolder1 := Folder{NameFolder: "SubOfSubFolder1"}

	SubOfSubFolder1.Add(&File{
		Name: "file3",
		Data: "hello file3",
	})

	subFolder1.Add(&SubOfSubFolder1)

	folder1.Add(&file1)
	folder1.Add(&file2)
	folder1.Add(&subFolder1)

	folder1.Search("datadata")
}

/*
		folder1
		/   |   \
	  data test subFolder1
	  				\
					subOfSubFolder1
						\
						file3

*/
