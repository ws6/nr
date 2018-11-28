package nr

//Given a two-dimensional table tab[1..i][1..j], such that tab[k][l] contains the number
//of events falling in bin k of one variable and bin l of another, this program returns Kendall’s τ
//as tau, its number of standard deviations from zero as z, and its two-sided significance level as
//prob. Small values of prob indicate a significant correlation (tau positive) or anticorrelation
//(tau negative) between the two variables. Although tab is a float array, it will normally
//contain integral values.
func Kendl2(tab [][]float64) (
	tau float64,
	z float64,
	prob float64,
	err error,
) {

	i, j := matrix_dim(tab)

	nn := i * j
	points := tab[i-1][j-1]
	en1 := 0.
	en2 := 0.
	s := 0.
	for k := 0; k <= nn-2; k++ {
		ki := k / j
		kj := k - j*ki
		points += tab[ki+1][kj+1]
		for l := k + 1; l <= nn-1; l++ {
			li := l / j
			lj := l - j*li
			m1 := li - ki
			m2 := lj - kj
			mm := m1 * m2
			pairs := tab[ki+1][kj+1] * tab[li+1][lj+1]

			if mm != 0 {
				en1 += pairs
				en2 += pairs

				if mm > 0 {
					s += pairs
				}
				if mm <= 0 {
					s += -pairs
				}
				continue
			}

			if m1 != 0 {
				en1 += pairs
			}
			if m2 != 0 {
				en2 += pairs
			}
		}

	}
	tau = s / sqrt(en1*en2)
	svar := (4.0*points + 10.0) / (9.0 * points * (points - 1.0))
	z = tau / sqrt(svar)
	prob = Erfcc(fabs(z) / 1.4142136)
	return
}
