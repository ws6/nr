package nr

func Lubksb(a [][]float64, n int, indx []int, b []float64) {
	sum := 0.
	ii := 0
	for i := 0; i < n; i++ {
		ip := indx[i]
		sum = b[ip]
		b[ip] = b[i]
		if ii != 0 {
			for j := ii; j <= i-1; j++ {
				sum -= a[i][j] * b[j]
			}
		}
		if ii == 0 {
			if sum != 0 {
				ii = i
			}
		}

		b[i] = sum
	}

	for i := n - 1; i >= 0; i-- {
		sum = b[i]
		for j := i + 1; j < n; j++ {
			sum -= a[i][j] * b[j]
		}
		b[i] = sum / a[i][i]
	}
}
