package nr

func Realft(data []float64, isign int) (
	err error,
) {

	n := len(data)
	if !ispower2(n) {
		err = ERR_NOT_POW_2
		return
	}
	theta := 3.141592653589793238 / float64(n>>1)
	c2 := 0.5
	c1 := 0.1
	if isign == 1 {
		c2 = -0.5
		if _err := Four1(data, 1); _err != nil {
			err = _err
			return
		}
	}
	if isign != 1 {
		theta = -theta
	}

	wtemp := sin(0.5 * theta)
	wpr := -2.0 * wtemp * wtemp
	wpi := sin(theta)
	wr := 1.0 + wpr
	wi := wpi

	for i := 1; i < (n >> 2); i++ {
		i1 := i + i
		i2 := 1 + (i1)
		i3 := n - i1
		i4 := 1 + (i3)
		h1r := c1 * (data[i1] + data[i3])
		h1i := c1 * (data[i2] - data[i4])
		h2r := -c2 * (data[i2] + data[i4])
		h2i := c2 * (data[i1] - data[i3])
		data[i1] = h1r + wr*h2r - wi*h2i
		data[i2] = h1i + wr*h2i + wi*h2r
		data[i3] = h1r - wr*h2r + wi*h2i
		data[i4] = -h1i + wr*h2i + wi*h2r
		wtemp := wr
		wr = (wtemp)*wpr - wi*wpi + wr
		wi = wi*wpr + wtemp*wpi + wi
	}
	h1r := data[0]
	if isign == 1 {

		data[0] = (h1r) + data[1]
		data[1] = h1r - data[1]
		return
	}
	data[0] = c1 * ((h1r) + data[1])
	data[1] = c1 * (h1r - data[1])
	if _err := Four1(data, -1); _err != nil {
		err = _err
		return
	}
	return
}
