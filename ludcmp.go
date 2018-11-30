package nr

//Given a matrix a[1..n][1..n], this routine replaces it by the LU decomposition of a rowwise
//permutation of itself. a and n are input. a is output, arranged as in equation (2.3.14) above;
//indx[1..n] is an output vector that records the row permutation effected by the partial
//pivoting; d is output as ±1 depending on whether the number of row interchanges was even
//or odd, respectively.
func Ludcmp(a [][]float64) (
	indx []int,
	d float64,
	err error,
) {
	n := len(a)
	_row, _col := matrix_dim(a)
	if n != _col || n != _row {
		err = ERR_NOT_SQURE_MATRIX
		return
	}
	indx = make([]int, n)
	d = 1.0

	vv := vector(n)
	for i := 0; i < n; i++ {
		big := 0.0
		for j := 0; j < n; j++ {
			temp := fabs(a[i][j])
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
			sum := a[i][j]
			for k := 0; k < i; k++ {
				sum -= a[i][k] * a[k][j]
			}
			a[i][j] = sum
		}
		imax := j
		big := 0.
		for i := j; i < n; i++ {
			sum := a[i][j]
			for k := 0; k < j; k++ {
				sum -= a[i][k] * a[k][j]
			}

			a[i][j] = sum
			dum := vv[i] * fabs(sum)
			if dum >= big {
				big = dum
				imax = i
			}
		}

		if j != imax {
			for k := 0; k < n; k++ {
				dum := a[imax][k]
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
		if j != (n - 1) {
			dum := 1.0 / (a[j][j])

			for i := j + 1; i < n; i++ {
				a[i][j] *= dum
			}
		}
	}

	return
}

func _Ludcmp(a [][]float64) (
	indx []int,
	d float64,
	err error,
) {
	n := len(a)
	_row, _col := matrix_dim(a)
	if n != _col || n != _row {
		err = nerror(`a is not a squred matrix, expect n x n`)
		return
	}
	indx = make([]int, n)

	for i := range indx {
		indx[i] = i
	}

	for i := 0; i < n; i++ {
		maxA := 0.
		imax := i

		for k := i; k < n; k++ {
			absA := fabs(a[k][i])
			if absA > maxA {
				maxA = absA
				imax = k
			}
		}

		if maxA < TINY {
			err = nerror(`failure, matrix is degenerate`)
			return
		}

		if imax != i {
			t := indx[i]
			indx[i] = indx[imax]
			indx[imax] = t

			//pivoting rows of A
			ptr := a[i]
			a[i] = a[imax]
			a[imax] = ptr
			//counting pivots starting from N (for determinant)
			indx[n]++
		}

		for j := i + 1; j < n; j++ {
			a[j][i] /= a[i][i]
			for k := i + 1; k < n; k++ {
				a[j][k] -= a[j][i] * a[i][k]
			}
		}

	}

	return
}
