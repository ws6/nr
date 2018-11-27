package nr

//Given the arrays bins1[1..nbins] and bins2[1..nbins], containing two sets of binned
//data, and given the number of constraints knstrn (normally 1 or 0), this routine returns the
//number of degrees of freedom df, the chi-square chsq, and the significance prob. A small value
//of prob indicates a significant difference between the distributions bins1 and bins2. Notethat
//bins1 and bins2 are both float arrays, although they will normally contain integer values.

func Chstwo(bins1, bins2 []float64, knstrn int) (
	df float64, //number of degrees of freedom
	chsq float64, //(nontrivially) the chi-square chsq
	prob float64, //significance
	err error,
) {
	n := len(bins1)
	_nbins := float64(n)
	if n != len(bins2) {
		err = nerror(`bins2 anre bins1 shall be same size`)
		return
	}
	df = _nbins - float64(knstrn)

	chsq = 0.

	for i := 0; i < n; i++ {
		if bins1[i] == 0. && bins2[i] == 0. {
			df -= 1
			continue
		}

		temp := bins1[i] - bins2[i]

		chsq += temp * temp / (bins1[i] + bins2[i])

	}

	prob, err = Gammq(0.5*df, 0.5*chsq)

	return
}
