package nr

//Returns the incomplete gamma function Q(a, x) evaluated by its continued fraction representation
//as gammcf. Also returns lnΓ(a) as gln.
func Gcf(a, x float64) (
	gammcf float64,
	gln float64,
	err error,
) {

	gln = Gammln(a)
	b := x + 1. - a
	c := 1. / FPMIN
	d := 1. / b
	h := d

	for i := 1; i <= MAXIT; i++ {
		an := float64(-i) * (float64(i) - a)
		b += 2.
		d = an*d + b
		if fabs(d) < FPMIN {
			d = FPMIN
		}
		c = b + an/c
		if fabs(c) < FPMIN {
			c = FPMIN
		}

		d = 1. / d
		del := d * c
		h *= del

		if fabs(del-1.) < EPS {
			gammcf = exp(-x+a*log(x)-gln) * h
			return
		}

	}

	err = nerror(`a too large, ITMAX too small in gcf`)

	return

}
