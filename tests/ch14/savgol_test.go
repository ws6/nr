package ch14

import (
	"math"
	"testing"

	r "github.com/ws6/nr"
)

func toBookPrecision(f float64) float64 {
	return math.Ceil(f*1000) / 1000.
}
func equalf(f1, f2, tolerance float64) bool {
	diff := f1 - f2
	if diff < 0 {
		diff = -diff
	}
	return diff < tolerance
}

func arrayEqualf(f1, f2 []float64) bool {
	n := len(f1)
	if n != len(f2) {
		return false
	}
	for i, v := range f1 {
		if !equalf(v, f2[i], 0.001) {
			return false
		}
	}

	return true

}
func TestSavgol(t *testing.T) {

	tests := [][]int{
		{2, 2, 2}, //m, nl, nr
		{2, 3, 1},
		{2, 4, 0},
		{2, 5, 5},
		{4, 4, 4},
		{4, 5, 5},
	}
	answers := [][]float64{ //from book table
		{-0.086, 0.343, 0.486, 0.343, -0.086},
		{-0.143, 0.171, 0.343, 0.371, 0.257},
		{0.086, -0.143, -0.086, 0.257, 0.886},
		{-0.084, 0.021, 0.103, 0.161, 0.196, 0.207, 0.196, 0.161, 0.103, 0.021, -0.084},
		{0.035, -0.128, 0.070, 0.315, 0.417, 0.315, 0.070, -0.128, 0.035},
		{0.042, -0.105, -0.023, 0.140, 0.280, 0.333, 0.280, 0.140, -0.023, -0.105, 0.042},
	}
	//Savgol(np, nl, nr, ld, m
	for i, test := range tests {

		nl := test[1]
		nr := test[2]
		m := test[0]
		ld := 0

		c, err := r.Savgol(nl, nr, ld, m)
		if err != nil {
			t.Fatal(err.Error())
		}
		sum := 0.
		for _, v := range c {
			sum += v
		}
		if !equalf(sum, 1., 0.001) {
			t.Fatal(`sum is not close to 1.`)
		}

		sorted := r.SortSavgolCoeff(c, nl, nr)
		ans := answers[i]
		if !arrayEqualf(ans, sorted) {
			t.Fail()
		}
	}

}
