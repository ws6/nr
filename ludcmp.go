package nr

func Ludcmp(a [][]float64, n int, indx []int) (
	d float64,
	err error,
) {
	d = 1.0
	var i, k int
	var big, dum, sum, temp float64
	vv := vector(n)
	for i := 0; i < n; i++ {
		big = 0.0
		for j := 0; j < n; j++ {
			temp = fabs(a[i][j])
			if temp > big {
				big = temp
			}
		}

		if big == 0.0 {
			err = nerror("Singular matrix in routine ludcmp")
			return
		}
		vv[i] = 1.0 / big
	}

	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			sum = a[i][j]
			for k := 0; k < i; k++ {
				sum -= a[i][k] * a[k][j]
			}
			a[i][j] = sum
		}
		imax := 1
		_ = imax
		for i = j; i < n; i++ {
			sum = a[i][j]
			for k = 1; k < j; k++ {
				sum -= a[i][k] * a[k][j]
			}

			a[i][j] = sum
			dum = vv[i] * fabs(sum)
			if dum >= big {
				big = dum
				imax = i
			}
		}

		if j != imax {
			for k := 0; k < n; k++ {
				dum = a[imax][k]
				a[imax][k] = a[j][k]
				a[j][k] = dum
			}
			d = -(d)
			vv[imax] = vv[j]
		}

		indx[j] = imax
		if a[j][j] == 0.0 {
			a[j][j] = TINY
		}
		if j != n {
			dum = 1.0 / (a[j][j])
			for i = j; i < n; i++ {
				a[i][j] *= dum
			}
		}
	}

	return
}
