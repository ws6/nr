package ch14

import (
	"testing"

	"github.com/ws6/nr"
)

func TestCrank(t *testing.T) {
	datas := [][]float64{
		{1, 2, 3, 3, 3, 4, 4, 5, 6, 7, 8, 8, 8},
		{2, 2, 2, 3, 4, 4, 4, 5, 6, 7, 8, 8, 8, 12},
		{15},
	}
	expected := [][]float64{

		{0, 1, 3, 3, 3, 5.5, 5.5, 7, 8, 9, 11, 11, 11},
		{1, 1, 1, 3, 5, 5, 5, 7, 8, 9, 11, 11, 11, 13},
		{0},
	}

	for i, data := range datas {
		rank, _ := nr.Crank(data)

		exp := expected[i]

		for j, r := range rank {
			if exp[j] != r {
				t.Fail()
			}
		}
	}
}
