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

func FMAX(f1, f2 float64) float64 {
	if f2 > f1 {
		return f2
	}

	return f1
}

func pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

func ipow(x, y int) float64 {
	return pow(float64(x), float64(y))
}

func IMIN(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func ispower2(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			break
		}
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}

func sin(f float64) float64 {
	return math.Sin(f)
}
