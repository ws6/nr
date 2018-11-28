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
	for _, data := range datas {
		rank, s := nr.Crank(data)
		t.Logf("%v %f", rank, s)
	}
}
