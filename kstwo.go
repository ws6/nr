package nr

import (
	"sort"
)

//Given an array data1[0..n1-1], and an array data2[0..n2-1], this routine returns the K–
//S statistic d, and the significance level prob for the null hypothesis that the data sets are
//drawn from the same distribution. Small values of prob showtha t the cumulative distribution
//function of data1 is significantly different from that of data2. The arrays data1 and data2
//are modified by being sorted into ascending order.

func Kstwo(data1, data2 []float64) (
	d float64, // K–S statistic d
	prob float64, //significance level prob
	err error,
) {
	n1 := len(data1)
	if n1 == 0 {
		err = nerror(`data1 is zero sized`)
		return
	}

	_n1 := float64(n1)

	n2 := len(data2)
	if n2 == 0 {
		err = nerror(`data2  is zero sized`)
		return
	}
	_n2 := float64(n2)

	fn1 := 0.
	fn2 := 0.
	j1 := 0
	j2 := 0
	sort.Float64s(data1)
	sort.Float64s(data1)
	d = 0.
	en1 := _n1
	en2 := _n2

	for {
		if !(j1 < n1 && j2 < n2) {
			break
		}
		d1 := data1[j1]
		d2 := data2[j2]

		if d1 <= d2 {
			fn1 = float64(j1) / en1
			j1++
		}
		if d2 <= d1 {
			fn2 = float64(j2) / en2
			j2++
		}

		dt := fabs(fn2 - fn1)
		if dt > d {
			d = dt
		}

	}
	en := sqrt(en1 * en2 / (en1 + en2))
	prob = Probks((en + 0.12 + 0.11/en) * d)
	return
}
