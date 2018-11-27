package nr

//Probks
//Kolmogorov-Smirnov probability function.
//Given an array data1[1..n1], and an array data2[1..n2], this routine returns the K–
//S statistic d, and the significance level prob for the null hypothesis that the data sets are
//drawn from the same distribution. Small values of prob showtha t the cumulative distribution
//function of data1 is significantly different from that of data2. The arrays data1 and data2
//are modified by being sorted into ascending order.
func Probks(alam float64) float64 {

	fac := 2.
	sum := 0.
	termbf := 0.

	a2 := -2. * alam * alam
	for j := 1; j <= 100; j++ {
		term := fac * exp(a2*float64(j*j))
		sum += term
		if fabs(term) <= EPS1*termbf || fabs(term) <= EPS2*sum {
			return sum
		}
		fac = -fac
		termbf = fabs(term)
	}

	return 1.
}
