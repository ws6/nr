package nr

//Two-dimensional Kolmogorov-Smirnov test on two samples. Given the x and y coordinates of
//the first sample as n1 values in arrays x1[1..n1] and y1[1..n1], and likewise for the second
//sample, n2 values in arrays x2 and y2, this routine returns the two-dimensional, two-sample
//K-S statistic as d, and its significance level as prob. Small values of prob show that the
//two samples are significantly different. Note that the test is slightly distribution-dependent, so
//prob is only an estimate.

func Ks2d2s(x1, y1, x2, y2 []float64) (
	d float64,
	prob float64,
	err error,
) {
	n1 := len(x1)
	if len(y2) != n1 {
		err = nerror(`x1 and y1 are not same size`)
		return
	}
	n2 := len(x2)
	if n2 != len(y2) {
		err = nerror(`x2 and y2 are not same size`)
		return
	}

	d1 := 0.

	for j := 0; j < n1; j++ {
		fa, fb, fc, fd, _err := Quadct(x1[j], y1[j], x1, y1)
		if _err != nil {
			err = _err
			return
		}

		ga, gb, gc, gd, _err := Quadct(x1[j], y1[j], x2, y2)
		if _err != nil {
			err = _err
			return
		}
		d1 = FMAX(d1, fabs(fa-ga))
		d1 = FMAX(d1, fabs(fb-gb))
		d1 = FMAX(d1, fabs(fc-gc))
		d1 = FMAX(d1, fabs(fd-gd))
	}
	d2 := 0.0
	for j := 0; j < n2; j++ {
		fa, fb, fc, fd, _err := Quadct(x2[j], y2[j], x1, y1)
		if _err != nil {
			err = _err
			return
		}

		ga, gb, gc, gd, _err := Quadct(x2[j], y2[j], x2, y2)
		if _err != nil {
			err = _err
			return
		}
		d2 = FMAX(d2, fabs(fa-ga))
		d2 = FMAX(d2, fabs(fb-gb))
		d2 = FMAX(d2, fabs(fc-gc))
		d2 = FMAX(d2, fabs(fd-gd))
	}

	d = .5 * (d1 + d2)

	sqen := sqrt(float64(n1*n2) / float64(n1+n2))
	r1, _, _, _err := Pearsn(x1, y1)
	if _err != nil {
		err = _err
		return
	}
	r2, _, _, _err := Pearsn(x2, y2)
	if _err != nil {
		err = _err
		return
	}

	rr := sqrt(1.0 - 0.5*(r1*r1+r2*r2))
	prob = Probks(d * sqen / (1.0 + rr*(0.25-0.75/sqen)))
	return
}
