package main

import "testing"

// Test<Nama func>
func TestSum(t *testing.T) {
	t.Log("Testing Sum total 4 with parameter 1 and 3")
	var sum1 = Sum(1, 3)
	if sum1 != 4 {
		t.Fail()
	}

	t.Log("Testing Sum total 8 with parameter 5 and 3")
	var sum2 = Sum(5, 3)
	if sum2 != 8 {
		t.Fail()
	}

	t.Log("Testing Sum total -3 with parameter -1 and -2")
	var sum3 = Sum(-1, -2)
	if sum3 != -3 {
		t.Fail()
	}

	t.Log("Testing Sum total 0 with parameter 0 and 0")
	var sum4 = Sum(0, 0)
	if sum4 != 0 {
		t.Errorf("Error result")
		t.Fail()
	}

	t.Log("Testing Sum total 1_000_000_000_001 with parameter 1_000_000_000_000 and 1")
	var sum5 = Sum(1_000_000_000_000, 1)
	if sum5 != 1_000_000_000_001 {
		t.Fail()
	}
}
