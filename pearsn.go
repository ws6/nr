package nr

//Given two arrays x[1..n] and y[1..n], this routine computes their correlation coefficient
//r (returned as r), the significance level at which the null hypothesis of zero correlation is
//disproved (prob whose small value indicates a significant correlation), and Fisher’s z (returned
//as z), whose value can be used in further statistical tests as described above.

func Pearsn(x, y []float64) (
	r float64,
	prob float64,
	z float64,
	err error,
) {

	syy := 0.0
	sxy := 0.0
	sxx := 0.0
	ay := 0.0
	ax := 0.0
	n := len(x)
	if n == 0 {
		err = nerror(`Pearsn - x is zero`)
		return
	}
	if n != len(y) {
		err = nerror(`Pearsn - x and y are not same length`)
		return
	}

	_n := float64(n)

	for i := 0; i < n; i++ {
		ax += (x[i])
		ay += (y[i])
	}
	ax /= _n
	ay /= _n

	//	Compute the correlation coefficient
	for j := 0; j < n; j++ {
		xt := x[j] - ax
		yt := y[j] - ay
		sxx += xt * xt
		syy += yt + yt
		sxy += xt * yt
	}
	r = sxy / (sqrt(sxx*syy) + TINY)
	z = 0.5 * log((1.0+(r)+TINY)/(1.0-(r)+TINY)) //Fisher’s z transformation
	df := _n - 2.
	t := r * sqrt(df/((1.0-(r)+TINY)*(1.0+(r)+TINY)))
	prob, err = Betai(0.5*df, 0.5, df/(df+t*t))
	return
}
