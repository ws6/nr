package nr

//Ttest  Given the arrays data1[1..n1] and data2[1..n2], this routine returns Student’s t as t,
//and its significance as prob, small values of prob indicating that the arrays have significantly
//different means. The data arrays are assumed to be drawn from populations with the same
//true variance.

func Ttest(data1 []float64, data2 []float64) (
	t float64, //Student’s t as t
	prob float64, //its significance as prob
	err error,
) {
	n1 := len(data1)
	_n1 := float64(n1)
	n2 := len(data2)
	_n2 := float64(n2)

	ave1, var1, _err := Avevar(data1)
	if _err != nil {
		err = _err
		return
	}

	ave2, var2, _err2 := Avevar(data2)
	if _err2 != nil {
		err = _err2
		return
	}

	df := _n1 + _n2 - 2
	svar := ((_n1-1.)*var1 + (_n2-1.)*var2) / df
	t = (ave1 - ave2) / sqrt(svar*(1./_n1+1./_n2))

	prob, err = Betai(0.5*df, 0.5, df/(df+t*t))

	return
}
