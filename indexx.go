package nr

import (
	"sort"
)

type Slice struct {
	sort.Float64Slice
	idx []int
}

func (s Slice) Swap(i, j int) {
	s.Float64Slice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

//Indexes an array arr[0..n-1], i.e., outputs the array indx[0..n-1] such that arr[indx[j]] is
//in ascending order for j = 1, 2, . . . ,N. The input quantities n and arr are not changed.
func Indexx(_arr []float64) []int {

	arr := make([]float64, len(_arr))

	for i, f := range _arr { //!!! to align with not changing arr in
		arr[i] = f
	}

	return _Indexx(arr)
}

func _Indexx(arr []float64) []int {

	s := &Slice{Float64Slice: sort.Float64Slice(arr), idx: make([]int, len(arr))}
	for i := range s.idx {
		s.idx[i] = i
	}
	sort.Sort(s)
	return s.idx
}
