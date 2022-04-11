package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Data evaluation digunakan untuk menampilkan value dari `struct` yang sudah didefinisikan dalam Golang sintax.
// Untuk menampilkan value dari field yang ada pada struct ini bisa dilakukan dengan cara `{{ .FieldName }}`.

func main() {
	type User struct {
		FirstName string
		Age       int
		Address   string
	}
	u := User{FirstName: "Afis", Age: 30, Address: "Jakarta"} // set value dari struct user dan disimpan ke variable u
	// uStruct := User{"Rogu", 17}

	// initiate new template dengan nama "test"
	tmpl, err := template.New("test").Parse("Usia {{.FirstName}} saat ini adalah {{.Age}} tahun. Alamat {{.Address}}") // ".FirstName" dan ".Age" adalah field name yang ada di struct User
	if err != nil {
		panic(err)
	}

	// execute si template bundle tmpl.Execute
	var b bytes.Buffer
	err = tmpl.Execute(&b, u)
	// hasil execute akan ditampung si b
	if err != nil {
		panic(err)
	}

	val := b.Bytes()
	fmt.Println(string(val))
}
