package nr

//Tptest
//Given the paired arrays data1[1..n] and data2[1..n], this routine returns Student’s t for
//paired data as t, and its significance as prob, small values of prob indicating a significant
//difference of means.
func Tptest(data1, data2 []float64) (
	t float64, //Student’s t for paired data as t,
	prob float64, //significance as prob
	err error,
) {

	n1 := len(data1)
	_n1 := float64(n1)
	n2 := len(data2)

	if n1 != n2 {
		err = nerror(`data1 and data2 not same size`)
		return
	}
	n := n1
	_n := _n1
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
	cov := 0.
	for i := 0; i < n; i++ {
		cov += (data1[i] - ave1) * (data2[i] - ave2)
	}
	df := _n - 1
	cov /= df
	sd := sqrt((var1 + var2 - 2.*cov) / _n)
	t = (ave1 - ave2) / sd
	prob, err = Betai(0.5*df, .5, df/(df+t*t))
	return
}
