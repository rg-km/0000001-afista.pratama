package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var dummyJSON = `
{
	"name" : "afis",
	"age" : 23,"address" : "lumajang / jember",
	"tinggi" : 172.32,
	"jomblo" : true,
	"mobil" : null
}
`

type DummyJSON struct {
	Name    string      `json:"name"`
	Age     int         `json:"age"`
	Address string      `json:"address"`
	Tinggi  float32     `json:"tinggi"`
	Jomblo  bool        `json:"jomblo"`
	Mobil   interface{} `json:"mobil"`
}

func main() {
	// json Marshal : dari struct / object golang ke json
	// json.Unmarshal: nangkap value dari json ke struct / object golang

	// var dummyData DummyJSON

	// // parameter 1 : casting ke []byte data jsonnya
	// // parameter 2 : pass reference / memori dari variable penampung struct
	// err := json.Unmarshal([]byte(dummyJSON), &dummyData)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("name dari json :", dummyData.Name)
	// 	fmt.Println("age dari json :", dummyData.Age)
	// 	fmt.Println("address dari json :", dummyData.Address)
	// 	fmt.Println("tinggi dari json :", dummyData.Tinggi)
	// }

	dataDummy := DummyJSON{
		Name:    "afis",
		Age:     23,
		Address: "jember",
		Tinggi:  172.32,
		Jomblo:  true,
	}

	respByte, err := json.Marshal(dataDummy)
	if err != nil {
		log.Panic(err)
	}

	// casting ke string
	fmt.Println(string(respByte))

	// package / module lain untuk mengirim ke luar output
	// I/O

	// http biasa
	// kirim ke http, tcp
	// simpan ke service lain
	// kirim ke bahasa lain
}
