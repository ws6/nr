package nr

import (
	"fmt"
)

//Given two real input arrays data1[1..n] and data2[1..n], this routine calls four1 and
//returns two complex output arrays, fft1[1..2n] and fft2[1..2n], each of complex length
//n (i.e., real length 2*n), which contain the discrete Fourier transforms of the respective data
//arrays. n MUST be an integer power of 2.
var ERR_NOT_SAME_SIZE = fmt.Errorf(`two array are not same size`)

func Twofft(data1, data2 []float64) (
	fft1 []float64,
	fft2 []float64,
	err error,
) {
	n := len(data1)
	if n != len(data1) {
		err = ERR_NOT_SAME_SIZE
		return
	}
	if !ispower2(n) {
		err = ERR_NOT_POW_2
		return
	}
	nn2 := n + n
	nn3 := 1 + (nn2)
	fft1 = make([]float64, n*2)
	fft2 = make([]float64, n*2)

	for j, jj := 0, 0; j < n; j, jj = j+1, jj+2 {
		fft1[jj] = data1[j]
		fft1[jj+1] = data2[j]
	}

	if _err := Four1(fft1, 1); _err != nil {
		err = _err
		return
	}
	fft2[0] = fft1[1]
	fft1[1] = 0.
	fft2[1] = 0.
	for j := 2; j < n+1; j += 2 {
		rep := 0.5 * (fft1[j] + fft1[nn2-j])
		rem := 0.5 * (fft1[j] - fft1[nn2-j])
		aip := 0.5 * (fft1[j+1] + fft1[nn3-j])
		aim := 0.5 * (fft1[j+1] - fft1[nn3-j])
		fft1[j] = rep
		fft1[j+1] = aim
		fft1[nn2-j] = rep
		fft1[nn3-j] = -aim
		fft2[j] = aip
		fft2[j+1] = -rem
		fft2[nn2-j] = aip
		fft2[nn3-j] = rem
	}
	return
}
