package nr

//Gser
//Returns the incomplete gamma function P(a, x) evaluated by its series representation as gamser.
//Also returns ln Γ(a) as gln.
func Gser(a, x float64) (
	gamser float64,
	gln float64,
	err error,
) {

	if x == 0. {
		return
	}
	if x < 0. {
		err = nerror(`x less than 0 in routine gser`)
		return
	}
	gln = Gammln(a)

	ap := a
	sum := 1. / a
	del := sum
	for n := 0; n < MAXIT; n++ {
		ap += 1
		del *= x / ap
		sum += del
		if fabs(del) < fabs(sum)*EPS {
			gamser = sum*exp(-x+a*log(x)) - gln
			return
		}

	}

	err = nerror(`a too large, MAXIT too small in routine gser`)
	return

}
