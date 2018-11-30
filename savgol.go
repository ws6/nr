package nr

//Returns in c[1..np], in wrap-around order (N.B.!) consistent with the argument respns in
//routine convlv, a set of Savitzky-Golay filter coefficients. nl is the number of leftward (past)
//data points used, while nr is the number of rightward (future) data points, making the total
//number of data points used nl+nr+1. ld is the order of the derivative desired (e.g., ld = 0
//for smoothed function). m is the order of the smoothing polynomial, also equal to the highest
//conserved moment; usual values are m = 2 or m = 4.

func Savgol(nl, nr, ld, m int) (
	c []float64,
	err error,
) {
	np := nl + nr + 1
	if np < nl+nr+1 || nl < 0 || nr < 0 || ld > m || nl+nr < m {
		err = nerror(`bad args in savgol`)
		return
	}

	a := matrix(m+1, m+1)
	for ipj := 0; ipj <= (m << 1); ipj++ {
		sum := 1.
		if ipj != 0 {
			sum = 0.
		}
		for k := 1; k <= nr; k++ {
			sum += pow(float64(k), float64(ipj))
		}
		for k := 1; k <= nl; k++ {
			sum += pow(float64(-k), float64(ipj))
		}
		mm := IMIN(ipj, 2*m-ipj)
		for imj := -mm; imj <= mm; imj += 2 {
			a[(ipj+imj)/2][(ipj-imj)/2] = sum
		}

	}

	indx, _, _err := Ludcmp(a)
	if _err != nil {
		err = _err
		return
	}

	b := vector(m + 1)
	for j := 0; j < m+1; j++ {
		b[j] = 0.0
	}
	b[ld] = 1.

	if _err = Lubksb(a, indx, b); _err != nil {
		err = _err
		return
	}

	c = make([]float64, np)
	for kk := 0; kk < np; kk++ {
		c[kk] = 0.0
	}

	//	info("b (after) = %v", b)
	for k := -nl; k <= nr; k++ {
		sum := b[0]
		for mm := 1; mm <= m; mm++ {
			sum += b[mm] * pow(float64(k), float64(mm)) //!!! book has a mistake here
		}
		kk := (np - k) % np
		c[kk] = sum
	}
	return
}

//SortSavgolCoeff
//sort the order as in the book
func SortSavgolCoeff(c []float64, nl, nr int) []float64 {
	np := nl + nr + 1

	ret := make([]float64, (np))
	i := 0
	for j := nl; j >= 0; j-- {

		ret[i] = c[j]
		i++

	}
	for j := 0; j < nr; j++ {
		ret[i] = c[np-1-j]
		i++

	}
	return ret
}
