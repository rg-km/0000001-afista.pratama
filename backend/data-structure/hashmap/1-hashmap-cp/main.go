package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}

	// 21, 28, 22
	// 21:2 , 22: 1, 28: 2

	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {
	// return map[int]int{}
	// map[int]int => umur, total yang memiliki umur sesuai key
	// loop
	// cek key di map => map[key]
	// 1. kalau belum ada, tambahkan key, dan value 1
	// 2. kalau key udah ada, increment, ++
	ages := make(map[int]int)
	for _, person := range people {
		ages[person.age]++

	}
	return ages
}

func FilterByAge(people []Person, age int) []Person {
	//return []Person{}
	// buat penampung []Person

	// lakukan looping
	// check tiap loopingan , kalau Person[index].age == age
	// kita append ke penampung seluruh datanya di index tersebut
	// return penampung
	var result []Person
	for _, person := range people {
		if person.age == age {
			result = append(result, person)
		}
	}
	return result
}
