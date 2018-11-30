package nr

const NR_END = 1

func matrix(nrow, ncol int) [][]float64 {
	ret := make([][]float64, nrow)

	for i := range ret {
		ret[i] = make([]float64, ncol)
	}
	return ret
}

func matrix_dim(d [][]float64) (row int, col int) {
	if len(d) == 0 {
		return
	}

	row = len(d)
	col = len(d[0])
	return
}

func vector(sz int) []float64 {
	return make([]float64, sz)
}
func ivector(sz int) []int {
	return make([]int, sz)
}

func imatrix(nrow, ncol int) [][]int {
	ret := make([][]int, nrow)

	for i := range ret {
		ret[i] = make([]int, ncol)
	}
	return ret
}

//imatrix_dim get row and col
func imatrix_dim(d [][]int) (row int, col int) {
	if len(d) == 0 {
		return
	}

	row = len(d)
	col = len(d[0])
	return
}
