package compareFast_test

import (
	"strconv"
	"testing"

	"github.com/ruang-guru/playground/backend/benchmarking-and-compiler/compare"
)

// after you complete the answer of this function
func BenchmarkUnmarshal(b *testing.B) {
	// use strconv.Itoa(n) to convert n to string, use UnmarshallFast here
	for n := 0; n < b.N; n++ {
		compare.UnmarshallFast(strconv.Itoa(n))
	}
}
