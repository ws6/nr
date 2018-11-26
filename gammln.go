package nr

var __cof__ = []float64{
	76.18009172947146, -86.50532032941677,
	24.01409824083091, -1.231739572450155,
	0.1208650973866179e-2, -0.5395239384953e-5,
}

func Gammln(xx float64) float64 {

	y := xx
	x := xx
	tmp := x + 5.5
	tmp -= (x + .5) * log(tmp) //!!!log10 or log?
	ser := 1.000000000190015
	for j := 0; j < len(__cof__); j++ {
		y += 1.
		ser += __cof__[j] / y
	}

	return -tmp + log(2.5066282746310005*ser/x)
}
