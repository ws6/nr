package nr

//Returns the complementary error function erfc(x) with fractional error everywhere less than
//1.2 × 10−7.

func Erfcc(x float64) float64 {
	z := fabs(x)
	t := 1. / (1. + .5*z)
	ans := t * exp(-z*z-1.26551223+t*(1.00002368+t*(0.37409196+t*(0.09678418+
		t*(-0.18628806+t*(0.27886807+t*(-1.13520398+t*(1.48851587+
			t*(-0.82215223+t*0.17087277)))))))))
	if x > 0. {
		return ans
	}

	return 2. - ans
}
