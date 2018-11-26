package nr

func Betai(a, b, x float64) (
	ret float64, // the incomplete beta function Ix(a, b).
	err error,
) {

	if x < 0. || x > 1. {
		err = nerror(`Bad x in routine betai`)
		return
	}

	bt := exp(Gammln(a+b) - Gammln(a) - Gammln(b) + a*log(x) + b*log(1.0-x))
	if x == 0.0 || x == 1.0 {
		bt = 0.0
	}

	if x < (a+1.0)/(a+b+2.0) {
		_betacf, _err := Betacf(a, b, x)
		if err != nil {
			err = _err
			return
		}
		ret = bt * _betacf / a

		return

	}
	_betacf, _err := Betacf(b, a, 1.0-x)
	if err != nil {
		err = _err
		return
	}
	ret = 1.0 - bt*_betacf/b
	return
}
