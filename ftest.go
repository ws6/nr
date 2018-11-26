package nr

//Ftest
//Given the arrays data1[1..n1] and data2[1..n2], this routine returns the value of f, and
//its significance as prob. Small values of prob indicate that the two arrays have significantly
//different variances.
func Ftest(data1, data2 []float64) (
	f float64,
	prob float64,
	err error,

) {
	n1 := len(data1)
	_n1 := float64(n1)
	n2 := len(data2)
	_n2 := float64(n2)

	_, var1, _err := Avevar(data1)
	if _err != nil {
		err = _err
		return
	}

	_, var2, _err2 := Avevar(data2)
	if _err2 != nil {
		err = _err2
		return
	}

	f = var2 / var1
	df1 := _n2 - 1
	df2 := _n1 - 1

	if var1 > var2 {
		f = var1 / var2
		df1 = _n1 - 1
		df2 = _n2 - 1
	}

	prob, err = Betai(.5*df2, .5*df1, df2/(df2+df1*f))
	if err != nil {
		return
	}
	prob = 2 * prob
	if prob > 1. {
		prob = 2. - prob
	}
	return
}
