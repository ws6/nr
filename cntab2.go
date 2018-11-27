package nr

//Given a two-dimensional contingency table in the form of an integer array nn[i][j], where i
//labels the x variable and ranges from 0 to ni-1, j labels the y variable and ranges from 0 to nj-1,
//this routine returns the entropy h of the whole table, the entropy hx of the x distribution, the
//entropy hy of the y distribution, the entropy hygx of y given x, the entropy hxgy of x given y,
//the dependency uygx of y on x (eq. 14.4.15), the dependency uxgy of x on y (eq. 14.4.16),
//and the symmetrical dependency uxy (eq. 14.4.17).

func Cntab2(nn [][]int) (
	h float64, //entropy h of the whole table
	hx float64, //the entropy hx of the x distribution
	hy float64, //the entropy hy of the y distribution
	hygx float64, //the entropy hygx of y given x,
	hxgy float64, // the entropy hxgy of x given y

	uygx float64, //the dependency uygx of y on x (eq. 14.4.15),
	uxgy float64, //the dependency uxgy of x on y (eq. 14.4.16),
	uxy float64, //the symmetrical dependency uxy (eq. 14.4.17).
	err error,
) {
	ni, nj := imatrix_dim(nn) //get row and col

	sumi := vector(ni)
	sumj := vector(nj)
	sum := 0.

	for i := 0; i < ni; i++ {
		sumi[i] = 0.
		for j := 0; j < nj; j++ {
			sumi[i] += float64(nn[i][j])
			sum += float64(nn[i][j])
		}

	}
	if sum == 0 {
		err = nerror(`sum is zero in cntab1`)
		return
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

	}

	hx = 0.

	for i := 0; i < ni; i++ {
		if sumi[i] == 0. {
			continue
		}
		p := sumi[i] / sum
		hx -= p * log(p)
	}
	hy = 0.
	for j := 0; j < nj; j++ {
		if sumj[j] == 0. {
			continue
		}
		p := sumj[j] / sum
		hy -= p * log(p)
	}
	h = 0.

	for i := 0; i < ni; i++ {
		for j := 0; j < nj; j++ {
			if nn[i][j] == 0. {
				continue
			}

			p := float64(nn[i][j]) / sum
			h -= p * log(p)
		}

	}

	hygx = h - hx
	hxgy = h - hy
	uygx = (hy - hygx) / (hy + TINY)
	uxgy = (hx - hxgy) / (hx + TINY)
	uxy = 2.0 * (hx + hy - h) / (hx + hy + TINY)
	return
}
