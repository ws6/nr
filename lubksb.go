package nr

import (
	"fmt"
)

var ERR_NOT_SQURE_MATRIX = fmt.Errorf(`not square matrics`)

//Solves the set of n linear equations A·X = B. Here a[1..n][1..n] is input, not as the matrix
//A but rather as its LU decomposition, determined by the routine ludcmp. indx[1..n] is input
//as the permutation vector returned by ludcmp. b[1..n] is input as the right-hand side vector
//B, and returns with the solution vector X. a, n, and indx are not modified by this routine
//and can be left in place for successive calls with different right-hand sides b. This routine takes
//into account the possibility that b will begin with many zero elements, so it is efficient for use
//in matrix inversion.
func Lubksb(a [][]float64, indx []int, b []float64) (
	err error,
) {

	n := len(a)
	_row, _col := matrix_dim(a)
	if n != _col || n != _row {
		err = ERR_NOT_SQURE_MATRIX
		return
	}

	if n != len(b) {
		err = nerror(`b is not same size as b`)
		return
	}

	if n != len(indx) {
		err = nerror(`indx is not same size as a`)
		return
	}

	ii := 0
	for i := 0; i < n; i++ {
		ip := indx[i]
		sum := b[ip]
		b[ip] = b[i]
		if ii != 0 {
			for j := ii - 1; j < i; j++ {
				sum -= a[i][j] * b[j]
			}
		} else if sum != 0 {

			ii = i + 1

		}

		b[i] = sum
	}

	for i := n - 1; i >= 0; i-- {
		sum := b[i]
		for j := i + 1; j < n; j++ {
			sum -= a[i][j] * b[j]
		}
		b[i] = sum / a[i][i]
	}

	return
}
