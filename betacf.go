package nr

const (
	MAXIT = 100     //Maximum allowed number of iterations
	EPS   = 3.0e-7  //Relative accuracy
	FPMIN = 1.0e-30 //Number near the smallest representable floating-point number
)

func Betacf(a, b, x float64) (h float64, err error) {
	qab := a + b
	qap := a + 1.
	qam := a - 1.

	c := 1.
	d := 1. - qab*x/qap
	if fabs(d) < FPMIN {
		d = FPMIN
	}
	d = 1. / d
	h = d
	m := 1
	for m := 1; m <= MAXIT; m++ {
		m2 := 2 * m
		_m := float64(m)
		_m2 := float64(m2)
		aa := _m * (b - _m) * x / ((qam + _m2) * (a + _m2))
		d = 1. + aa*d
		if fabs(d) < FPMIN {
			d = FPMIN
		}
		c = 1. + aa/c
		if fabs(c) < FPMIN {
			c = FPMIN
		}

		d = 1. / d
		h *= d * c
		aa = -(a + _m) * (qab + _m) * x / ((a + _m2) * (qap + _m2))
		d = 1.0 + aa*d
		if fabs(d) < FPMIN {
			d = FPMIN
		}
		if fabs(c) < FPMIN {
			c = FPMIN
		}
		d = 1. / d
		del := d * c
		h *= del
		if fabs(del-1.) < EPS {
			break
		}
	}

	if m > MAXIT {
		err = nerror(`a or b too big, or MAXIT too small in betacf`)
		return
	}

	return
}
