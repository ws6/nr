package nr

//Two-dimensional Kolmogorov-Smirnov test of one sample against a model. Given the x and y
//coordinates of n1 data points in arrays x1[1..n1] and y1[1..n1], and given a user-supplied
//function quadvl that exemplifies the model, this routine returns the two-dimensional K-S
//statistic as d1, and its significance level as prob. Small values of prob show that the sample
//is significantly different from the model. Note that the test is slightly distribution-dependent,
//so prob is only an estimate.

func Ks2d1s(x1, y1 []float64, quadvl func(float64, float64) (float64, float64, float64, float64)) (
	r float64,
	prob float64,
	z float64,
	err error,
) {
	n1 := len(x1)
	if n1 != len(y1) {
		err = nerror(`x1 and y1 are not same size`)
		return
	}
	d1 := 0.
	for j := 0; j < n1; j++ {
		fa, fb, fc, fd, _err := Quadct(x1[j], y1[j], x1, y1)
		if _err != nil {
			err = _err
			return
		}
		ga, gb, gc, gd := quadvl(x1[j], y1[j])

		d1 = FMAX(d1, fabs(fa-ga))
		d1 = FMAX(d1, fabs(fb-gb))
		d1 = FMAX(d1, fabs(fc-gc))
		d1 = FMAX(d1, fabs(fd-gd))
	}

	r1, _, _, _err := Pearsn(x1, y1)
	if _err != nil {
		err = _err
		return
	}
	sqen := sqrt(float64(n1))
	rr := sqrt(1. - r1*r1)
	prob = Probks(d1 * sqen / (1. + rr*(0.25-.75/sqen)))
	return
}
