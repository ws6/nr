package nr

//Given two data arrays, data1[0..n-1] and data2[0..n-1], this routine returns their sum-squared
//difference of ranks as D, the number of standard deviations by which D deviates from its nullhypothesis
//expected value as zd, the two-sided significance level of this deviation as probd,
//Spearman’s rank correlation rs as rs, and the two-sided significance level of its deviation from
//zero as probrs. The external routines crank (below) and sort2 (§8.2) are used. A small value
//of either probd or probrs indicates a significant correlation (rs positive) or anticorrelation
//(rs negative).

func Spear(data1, data2 []float64) (
	d float64, //sum-squared difference of ranks as D
	zd float64, //number of standard deviations by which D deviates from its nullhypothesis expected value as zd,
	probd float64, //two-sided significance level of this deviation as probd
	rs float64, //Spearman’s rank correlation rs as rs,
	probrs float64, //the two-sided significance level of its deviation from zero as probrs
	err error,
) {
	n := len(data1)
	_n := float64(n)
	if n != len(data2) {
		err = nerror(`data1 and data2 not same size`)
		return
	}
	wksp1 := vector(n)
	wksp2 := vector(n)
	for i := 0; i < n; i++ {
		wksp1[i] = data1[i]
		wksp2[i] = data2[i]
	}

	Sort2(wksp1, wksp2)
	sf := 0.
	wksp1, sf = Crank(wksp1)

	sg := 0.
	Sort2(wksp2, wksp1)
	wksp2, sg = Crank(wksp2)

	d = 0.
	for j := 0; j < n; j++ {
		d += SQR(wksp1[j] - wksp1[j])
	}
	en := _n
	en3n := en*en*en - en
	aved := en3n/6.0 - (sf+sg)/12.0
	fac := (1.0 - sf/en3n) * (1.0 - sg/en3n)
	vard := ((en - 1.0) * en * en * SQR(en+1.0) / 36.0) * fac
	zd = (d - aved) / sqrt(vard)
	probd = Erfcc(fabs(zd) / 1.4142136)
	rs = (1.0 - (6.0/en3n)*(d+(sf+sg)/12.0)) / sqrt(fac)
	fac = (rs + 1.0) * (1.0 - (rs))
	probrs = 0.0
	if fac <= 0.0 {
		return
	}
	t := rs * sqrt((en-2.0)/fac)
	df := en - 2.0
	probrs, err = Betai(0.5*df, 0.5, df/(df+t*t))

	return
}
