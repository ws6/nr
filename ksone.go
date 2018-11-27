package nr

import (
	"sort"
)

//Given an array data[1..n], and given a user-supplied function of a single variable func which
//is a cumulative distribution function ranging from 0 (for smallest values of its argument) to 1
//(for largest values of its argument), this routine returns the K–S statistic d, and the significance
//level prob. Small values of prob showtha t the cumulative distribution function of data is
//significantly different from func. The array data is modified by being sorted into ascending
//order.

func Ksone(data []float64, userfn func(float64) float64) (
	d float64, //K–S statistic d,
	prob float64, //significance level

) {
	fo := 0.
	n := len(data)
	_n := float64(n)
	sort.Float64s(data)

	en := _n
	d = 0.
	for j := 0; j < n; j++ {
		fn := float64(j+1) / en
		ff := userfn(data[j])

		dt := FMAX(fabs(fo-ff), fabs(fn-ff))

		if dt > d {
			d = dt
		}

		fo = fn

	}

	en = sqrt(en)
	prob = Probks((en + 0.12 + 0.11/en) * (d))
	return

}
