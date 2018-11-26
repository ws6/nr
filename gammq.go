package nr

func Gammq(a, x float64) (
	ret float64,
	err error,
) {
	if x < 0.0 || a <= 0.0 {
		err = nerror(`Invalid arguments in routine gammq`)
		return
	}
	if x < (a + 1.) {
		gamser, _, _err := Gser(a, x)
		if _err != nil {
			err = _err
			return
		}

		ret = 1. - gamser
		return
	}

	gammcf, _, _err := Gcf(a, x)

	if _err != nil {
		err = _err
		return
	}

	ret = gammcf

	return

}
