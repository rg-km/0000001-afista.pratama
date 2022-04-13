package main

import "fmt"

// tipe data set, data structure (data terapan)
// karakteristik
// - setiap element itu unik
// - tidak ada index, karena disusun acak / disusun tidak berurutan
// - mutable , datanya bisa diubah2 (bisa ditambah, bisa dihapus)

func main() {
	// [1,1,2,2,3,3]
	// {1,2,3} => tambah data 2 => error / skip

	// penerapan set
	// inisialisasi
	set := MakeSet()
	set2 := NewSet("afis", "jaya", "jember")

	// fmt.Println("set :", set.Size())
	// set.Print()

	// fmt.Println("set2 :", set2.Size())
	// set2.Print()

	// otak atik
	set.Add("afis")
	set.Add("afis")

	set2.Add("afis")
	set2.Add("afis")
	set2.Add("jaya")

	set.AddMultiply("1", "2", "3")

	set2.Remove("afis")

	set.Print()
	set2.Print()

}

// encapsulation
func MakeSet() *dataSet {
	return &dataSet{
		set: make(map[string]bool),
	}
}

// encapsulation dinamis / complex
// var1 = NewSet("1","2", "3", "4","5")
func NewSet(data ...string) *dataSet {
	var set = make(map[string]bool)

	for _, d := range data {
		set[d] = true
	}

	return &dataSet{
		set: set,
	}
}

// struct + map
type dataSet struct { // private data
	set map[string]bool
}

// required method Add, Remove, Exist, Size
func (s *dataSet) Exist(data string) bool {
	if _, ok := s.set[data]; ok {
		return true
	} else {
		return false
	}
}

func (s *dataSet) Add(data string) {
	// if check := s.Exist(data); check {
	// 	panic(errors.New("data exist in set"))
	// }

	s.set[data] = true
}

func (s *dataSet) AddMultiply(data ...string) {
	for _, d := range data {
		s.set[d] = true
	}
}

func (s *dataSet) Remove(data string) error {
	_, exists := s.set[data]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(s.set, data)
	return nil
}

func (s *dataSet) Size() int {
	return len(s.set)
}

func (s *dataSet) Print() {
	/*
		"afis" : true
		"jaya" : true
		"jember" : true

		["afis", "jaya", "jember"]
	*/

	var output []string

	for key, _ := range s.set {
		output = append(output, key)
	}

	fmt.Println(output)
}
