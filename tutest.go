package nr

//Tutest
//Given the arrays data1[1..n1] and data2[1..n2], this routine returns Student’s t as t, and
//its significance as prob, small values of prob indicating that the arrays have significantly different
//means. The data arrays are allowed to be drawn from populations with unequal variances
func Tutest(data1, data2 []float64) (
	t float64, //Student’s t
	prob float64, //significance as prob
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

	t = (ave1 - ave2) / sqrt(var1/_n1+var2/_n2)
	df := SQR(var1/_n1+var2/_n2) / (SQR(var1/_n1)/(_n1-1.) + SQR(var2/_n2)/(_n2-1.))
	prob, err = Betai(0.5*df, 0.5, df/(df+SQR(t)))
	return
}
