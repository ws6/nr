package nr

import (
	"math"
)

func fabs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

func log(f float64) float64 {
	return math.Log(f)
}
func exp(f float64) float64 {
	return math.Exp(f)
}

func sqrt(f float64) float64 {
	return math.Sqrt(f)
}

func SQR(f float64) float64 {
	return f * f
}
