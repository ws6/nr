package nr

//Given an origin (x, y), and an array of nn points with coordinates xx[1..nn] and yy[1..nn],
//count how many of them are in each quadrant around the origin, and return the normalized
//fractions. Quadrants are labeled alphabetically, counterclockwise from the upper right. Used
//by ks2d1s and ks2d2s.

func Quadct(x, y float64, xx, yy []float64) (
	fa float64,
	fb float64,
	fc float64,
	fd float64,
	err error,
) {
	nn := len(xx)
	if nn != len(yy) {
		err = nerror(`xx and yy are not same size`)
		return
	}
	na := 0
	nb := 0
	nc := 0
	nd := 0
	for k := 0; k < nn; k++ {
		if yy[k] > y {

			if xx[k] > x {
				na++
			}
			if xx[k] <= x {
				nb++
			}
			continue

		}
		if xx[k] > x {
			nd++
		}
		if xx[k] <= x {
			nc++
		}
	}

	ff := 1. / float64(nn)
	fa = ff * float64(na)
	fb = ff * float64(nb)
	fc = ff * float64(nc)
	fd = ff * float64(nd)
	return
}
