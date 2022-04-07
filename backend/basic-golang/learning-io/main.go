package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// stringer
// mengubah output

// type Cat struct {
// 	Name  string
// 	Sound string
// }

func ReadFile() []byte {
	filename := "input.txt"

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}

	return byteData
}

func SumAll(data []byte) int {
	numbers := strings.Split(string(data), " ")

	// change to int
	var ints []int
	for _, n := range numbers {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Panic(err)
		}

		ints = append(ints, num)
	}

	var result int
	for _, i := range ints {
		result += i
	}

	return result
}

func WriteFile(result int) {
	// buat file dulu
	file, err := os.Create("output.txt")
	if err != nil {
		log.Panic(err)
	}

	n, err := file.WriteString(fmt.Sprintf("hasil penjumlahan semua : %v", result))
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("result writeString", n)
}

func ReadFileByLine() {
	filename := "linebyline.txt"

	// 2 cara
	// file, err := os.Open(filename)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// defer file.Close()

	// reader := bufio.NewReader(file)

	// for {
	// 	line, _, err := reader.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	fmt.Println(string(line))
	// }

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(byteData))
}

func main() {
	// data := ReadFile()
	ReadFileByLine()

	// result := SumAll(data)

	// WriteFile(result)

	// cat1 := Cat{Name: "dorito", Sound: "meow"}

	// fmt.Println(cat1) // proses output {"dorito", "meow"}
}

// interface Stringer, pakai method String
// func (c Cat) String() string {
// 	return fmt.Sprintf("%v, my name is %v", c.Sound, c.Name)
// }
