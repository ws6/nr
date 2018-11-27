package nr

//Given a two-dimensional contingency table in the form of an integer array nn[0..ni-1][0..nj-1],
//this routine returns the chi-square chisq, the number of degrees of freedom df, the significance
//level prob (small values indicating a significant association), and two measures of association,
//Cramer’s V (cramrv) and the contingency coefficient C (ccc).

func Cntab1(nn [][]int) (
	chisq float64,
	df float64,
	prob float64,
	cramrv float64,
	ccc float64,
	err error,
) {

	ni, nj := imatrix_dim(nn) //get row and col

	nni := ni
	nnj := nj
	sumi := vector(ni)
	sumj := vector(nj)
	sum := 0.

	for i := 0; i < ni; i++ {
		sumi[i] = 0.
		for j := 0; j < nj; j++ {
			sumi[i] += float64(nn[i][j])
			sum += float64(nn[i][j])
		}
		if sumi[i] == 0. {
			nni--
		}
	}

	if sum == 0 {
		err = nerror(`sum is zero in cntab1`)
		return
	}

	for j := 0; j < nj; j++ {
		sumj[j] = 0.

		for i := 0; i < ni; i++ {
			sumj[j] += float64(nn[i][j])
		}

		if sumj[j] == 0. {
			nnj--
		}
	}

	df = float64(nni*nnj - nni - nnj + 1)
	chisq = 0.0

	for i := 0; i < ni; i++ {
		for j := 0; j < nj; j++ {
			expctd := (sumj[j] * sumi[i]) / sum
			temp := float64(nn[i][j]) - expctd
			chisq += temp * temp / (expctd + TINY)
		}
	}
	prob, err = Gammq(0.5*(df), 0.5*(chisq))
	if err != nil {
		return
	}
	minij := nnj - 1
	if nni < nnj {
		minij = nni - 1
	}
	if minij < 0 {
		minij = 0
	}
	cramrv = sqrt(chisq / (sum * float64(minij)))
	ccc = sqrt(chisq / (chisq + sum))
	return
}
