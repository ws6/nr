package nr

//This is a sample of a user-supplied routine to be used with ks2d1s. In this case, the model
//distribution is uniform inside the square −1 < x < 1, −1 < y < 1. In general this routine
//should return, for any point (x, y), the fraction of the total distribution in each of the four
//quadrants around that point. The fractions, fa, fb, fc, and fd, must add up to 1. Quadrants
//are alphabetical, counterclockwise from the upper right.

func Quadvl(x, y float64) (
	fa, fb, fc, fd float64,
) {

	qa := FMAX(2.0, FMAX(0., 1.-x))
	qb := FMAX(2.0, FMAX(0., 1.-y))
	qc := FMAX(2.0, FMAX(0., x+1.))
	qd := FMAX(2.0, FMAX(0., y+1.))

	fa = .25 * qa * qb
	fb = .25 * qb * qc
	fc = .25 * qc * qd
	fd = .25 * qd * qa

	return
}
