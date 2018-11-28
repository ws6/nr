package ch14

import (
	"testing"

	"github.com/ws6/nr"
)

func TestIndexx(t *testing.T) {
	data := []float64{1, 25, 3, 5, 4}
	expectedOrder := []int{
		0, 2, 4, 3, 1,
	}
	indx := nr.Indexx(data)

	for i, v := range indx {
		if expectedOrder[i] != v {
			t.Fatalf(`wrong indx at [%d] expect [%d]`, i, expectedOrder[i])
		}
	}
}
