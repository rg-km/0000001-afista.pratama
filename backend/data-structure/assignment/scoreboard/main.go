package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

// Less adalah interface , bisa kita isi logic di dalamnya
// ngecek value di index i dengan value di index j

func (s Scores) Less(i, j int) bool {
	// check pertama
	// jika total score di index i lebih besar dari total score di index J
	// kembalikan true
	// jika total score di index j lebih besar dari total score di index i
	// kembalikan false

	// Nilai = 4 x Jumlah Benar - 1 x Jumlah Salah

	scoreI := (s[i].Correct * 4) - (s[i].Wrong * 1)
	scoreJ := (s[j].Correct * 4) - (s[j].Wrong * 1)

	// case 1 , mengurutkan dari terbesar ke terkecil
	if scoreI > scoreJ {
		return true
	} else if scoreI < scoreJ {
		return false
	} else {
		// case 2: kalau score sama, kita bandingkan dengan correct terbanyak
		if s[i].Correct > s[j].Correct {
			return true
		} else if s[i].Correct < s[j].Correct {
			return false
		} else {
			// case 3 : score sama, total correct sama, cari nama terkecil/ urutan dari depan

			// string juga bisa dicompare
			// "aa" < "bb" , lebih besar "bb" >> binary
			// "a" > "aa" , lebih besar "aa" >> binary

			//"aa" "bb",
			// 0 atau 1, base 2 , angka 0 sama 1
			// 0 = 0
			// 1 = 1
			// 2 => 10
			// 3 => 11
			// 4 => 110
			return s[i].Name < s[j].Name
		}
	}
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 3, 4, 0},
		{"Doni", 3, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})
	sort.Sort(scores)

	fmt.Println(scores.TopStudents())
}
