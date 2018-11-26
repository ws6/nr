package nr

const NR_END = 1

func matrix(nrow, ncol int) [][]float64 {
	ret := make([][]float64, nrow)

	for i := range ret {
		ret[i] = make([]float64, ncol)
	}
	return ret
}

func vector(sz int) []float64 {
	return make([]float64, sz)
}
