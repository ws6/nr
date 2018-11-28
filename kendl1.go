package nr

//Given data arrays data1[1..n] and data2[1..n], this program returns Kendall’s τ as tau,
//its number of standard deviations from zero as z, and its two-sided significance level as prob.
//Small values of prob indicate a significant correlation (tau positive) or anticorrelation (tau
//negative).

func Kendl1(data1, data2 []float64) (
	tau float64,
	z float64,
	prob float64,
	err error,
) {
	n := len(data1)
	if n != len(data1) {
		err = nerror(`data1 and data2 not same size`)
		return
	}
	_n := float64(n)
	n1 := 0
	n2 := 0
	is := 0

	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			a1 := data1[j] - data1[k]
			a2 := data2[j] - data2[k]
			aa := a1 * a2
			if aa != 0 {
				n1++
				n2++
				if aa > 0 {
					is++
				}
				if aa <= 0 {
					is--
				}

				continue
			}

			if a1 != 0 {
				n1++
			}
			if a2 != 0 {
				n2++
			}
		}
	}

	tau = float64(is) / (sqrt(float64(n1)) * sqrt(float64(n2)))
	svar := (4.0*_n + 10.0) / (9.0 * _n * (_n - 1.0))
	z = tau / sqrt(svar)
	prob = Erfcc(fabs(z) / 1.4142136)
	return
}
