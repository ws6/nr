package nr

//Convolves or deconvolves a real data set data[1..n] (including any user-supplied zero padding)
//with a response function respns[1..n]. The response function must be stored in wrap-around
//order in the first m elements of respns, where m is an odd integer ≤ n. Wrap-around order
//means that the first half of the array respns contains the impulse response function at positive
//times, while the second half of the array contains the impulse response function at negative times,
//counting down from the highest element respns[m]. On input isign is +1 for convolution,
//−1 for deconvolution. The answer is returned in the first n components of ans. However,
//ans must be supplied in the calling program with dimensions [1..2*n], for consistency with
//twofft. n MUST be an integer power of two.

func Convlv(data, respns []float64, isign int) (
	ans []float64,
	err error,
) {

	n := len(data)
	if !ispower2(n) {
		err = ERR_NOT_POW_2
		return
	}
	ans = make([]float64, 2*n)
	m := len(respns)

	temp := make([]float64, n)

	temp[0] = respns[0]
	for i := 1; i < (m+1)/2; i++ {
		temp[i] = respns[i]
		temp[n-i] = respns[m-i]
	}

	for i := (m + 1) / 2; i < n-(m-1)/2; i++ {
		temp[i] = 0.0
	}

	for i := 0; i < n; i++ {
		ans[i] = data[i]
	}
	if _err := Realft(ans, 1); _err != nil {
		err = _err
		return
	}
	if _err := Realft(temp, 1); _err != nil {
		err = _err
		return
	}
	no2 := float64(n >> 1)
	if isign == 1 {
		for i := 2; i < n; i += 2 {
			tmp := ans[i]
			ans[i] = (ans[i]*temp[i] - ans[i+1]*temp[i+1]) / no2
			ans[i+1] = (ans[i+1]*temp[i] + tmp*temp[i+1]) / no2
		}
		ans[0] = ans[0] * temp[0] / no2
		ans[1] = ans[1] * temp[1] / no2
		if _err := Realft(ans, -1); _err != nil {
			err = _err
			return
		}
		return
	}
	if isign == -1 {
		for i := 2; i < n; i += 2 {
			mag2 := SQR(temp[i]) + SQR(temp[i+1])
			if mag2 == 0.0 {
				err = nerror("Deconvolving at response zero in convlv")
				return
			}
			tmp := ans[i]
			ans[i] = (ans[i]*temp[i] + ans[i+1]*temp[i+1]) / mag2 / no2
			ans[i+1] = (ans[i+1]*temp[i] - tmp*temp[i+1]) / mag2 / no2
		}
		if temp[0] == 0.0 || temp[1] == 0.0 {
			err = nerror("Deconvolving at response zero in convlv")
			return
		}

		ans[0] = ans[0] / temp[0] / no2
		ans[1] = ans[1] / temp[1] / no2

		if _err := Realft(ans, -1); _err != nil {
			err = _err
			return
		}
		return
	}
	err = nerror(`No meaning for isign in convlv`)
	return
}
