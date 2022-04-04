package main

import (
	"fmt"
	"log"
	"math" // package
)

// interface adalah tipe data

// method signature ?

type Circle struct {
	Radius int
	// 2 method , menghitung luas, menghitung keliling
}

// method
// kepemilikan
func (c Circle) Luas() float32 {
	return math.Pi * float32(c.Radius) * float32(c.Radius)
}

func (c Circle) Keliling() float32 {
	return 2 * math.Pi * float32(c.Radius)
}

type Square struct {
	Width  int
	Height int
	// 2 method, menghitung luas, menghitung keliling
}

func (s Square) Luas() float32 {
	return float32(s.Width) * float32(s.Height)
}

func (s Square) Keliling() float32 {
	return float32(2 * (s.Width + s.Height))
}

type Triangle struct {
	Length int
	Height int
	// 2 method
}

type Geometri interface {
	// signature
	Luas() float32
	Keliling() float32
}

// func ShowKelilingdanLuas(geo Geometri) {
// 	// apapun datanya, yang penting dia punya func / method Luas dan Keliling
// 	fmt.Println("keliling :", geo.Keliling())
// 	fmt.Println("luas :", geo.Luas())
// }

func ShowKelilingdanLuas(geo Geometri) {
	fmt.Println("keliling :", geo.Luas())
	fmt.Println("Luas :", geo.Luas())
}

// tidak boleh ada perbuahan tipe data
// harus banyak buat if else
func SumApapundengan2(apapun interface{}) int {
	if apapun.(string) == "afista" {
		log.Fatal("error tipe data bukan int")
	}
	return apapun.(int) + 2
}

func main() {
	// implementasi
	// tipe data yang bisa diisi apapun -> interface kosong
	// method signature

	// method signature
	// circle1 := Circle{Radius: 5}

	// square1 := Square{Width: 10, Height: 5}

	// triange1 := Triangle{Length: 10, Height: 15}

	SumApapundengan2("afista")

	// ShowKelilingdanLuas(circle1)
	// ShowKelilingdanLuas(square1)

	// ShowKelilingdanLuas(triange1)

	// interface kosong
	// default value = <nil> / null di bahasa lain
	// bisa diisi apapun
	// tidak bisa dioperasi dengan tipe lain, harus di casting
	// tidak bisa dioperasi matematic , walaupun sesama interface
	// var apapun interface{}

	// error
	// interface

	// pengecekannya menggunakan nil
	// if err != nil {
	// execute error
	// }
	// fmt.Println(apapun)

	// apapun = "afista"

	// apapun = 200
	// fmt.Println(apapun)
	// apapun = 3.14
	// fmt.Println(apapun)
	// apapun = nil
	// fmt.Println(apapun)
	// apapun = errors.New("error: error")
	// fmt.Println(apapun)

	// fmt.Println(apapun)

	// casting
	// namaLengkap := apapun.(string) + " " + "pratama"
	// var apapun2 interface{} = 100

	// operasi matematic = + - / * , tidak bisa diimplement ke interface
	// hasil := apapun.(string) + " pratama"
	// .(int)
	// .(string)
	// .(bool) => boolean (true / false)
	// .(float32) => 3.14
	// .(error) => menjadi error

	// fmt.Println(hasil)

}
