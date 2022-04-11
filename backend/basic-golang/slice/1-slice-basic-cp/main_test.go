package main

import "testing"

// function test => Test
func TestAddDataToSlice(t *testing.T) {
	// 3A, inisialisasi
	var res []string

	// ACT A2
	res = AddToSlice("Afista", "Pratama")

	// assertion
	if len(res) != 2 {
		t.Errorf("len must be 2")
	}

	// assertion
	if res[0] != "Afista" || res[1] != "Pratama" {
		t.Errorf("inserting error")
	}
}
