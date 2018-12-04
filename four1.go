package nr

import (
	"fmt"
)

var ERR_NOT_POW_2 = fmt.Errorf(`not power of 2`)

//Replaces data[1..2*nn] by its discrete Fourier transform, if isign is input as 1; or replaces
//data[1..2*nn] by nn times its inverse discrete Fourier transform, if isign is input as −1.
//data is a complex array of length nn or, equivalently, a real array of length 2*nn. nn MUST
//be an integer power of 2 (this     checked for!).
func Four1(data []float64, isign int) (
	err error,
) {
	n := len(data)
	if !ispower2(n) {
		err = ERR_NOT_POW_2
		return
	}
	nn := n / 2
	j := 1
	for i := 1; i < n; i += 2 {
		if j > i {
			data[j-1], data[i-1] = data[i-1], data[j-1]
			data[j], data[i] = data[i], data[j]
		}
		m := nn
		for m >= 2 && j > m {
			j -= m
			m >>= 1
		}

		j += m
	}
	mmax := 2
	for n > mmax {
		mmax <<= 1
		istep := mmax
		theta := float64(isign) * (6.28318530717959 / float64(mmax))
		wtemp := sin(0.5 * theta)
		wpr := -2.0 * wtemp * wtemp
		wpi := sin(theta)
		wr := 1.0
		wi := 0.

		for m := 1; m < mmax; m += 2 {
			for i := m; i <= n; i += istep {
				j = i + mmax
				tempr := wr*data[j-1] - wi*data[j]
				tempi := wr*data[j] + wi*data[j-1]
				data[j-1] = data[i-1] - tempr
				data[j] = data[i] - tempi
				data[i-1] += tempr
				data[i] += tempi
			}
			wtemp = wr
			wr = wtemp*wpr - wi*wpi + wr
			wi = wi*wpr + wtemp*wpi + wi
		}
		mmax = istep
	}
	return

}
