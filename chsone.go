package nr

//Chsone chi-square implementation 1
//Given the array bins[1..nbins] containing the observed numbers of events, and an array
//ebins[1..nbins] containing the expected numbers of events, and given the number of constraints
//knstrn (normally one), this routine returns (trivially) the number of degrees of freedom
//df, and (nontrivially) the chi-square chsq and the significance prob. A small value of prob
//indicates a significant difference between the distributions bins and ebins. Note that bins
//and ebins are both float arrays, although bins will normally contain integer values.

func Chsone(bins, ebins []float64, knstrn int) (
	df float64, //number of degrees of freedom
	chsq float64, //(nontrivially) the chi-square chsq
	prob float64, //significance
	err error,
) {
	n := len(bins)
	_nbins := float64(n)
	if n != len(ebins) {
		err = nerror(`bins anre ebins shall be same size`)
		return
	}

	df = _nbins - float64(knstrn)

	chsq = 0.

	for i := 0; i < n; i++ {
		if ebins[i] <= 0. {
			err = nerror(`  expected number shall be greater than 0 `)
			return
		}
		temp := bins[i] - ebins[i]
		chsq += temp * temp / ebins[i]
	}

	prob, err = Gammq(.5*df, .5*chsq)

	return
}
