package nr

var ERR_TWO_MORE = nerror(`must be at least 2 in data`)

func Moment(data []float64) (
	ave float64, //average
	adev float64, //average deviation
	sdev float64, //standard deviation
	variance float64, //variance
	skew float64, //skewness
	curt float64, // kurtosis
	err error,
) {
	n := len(data)
	fn := float64(n)
	if n <= 1 {
		err = ERR_TWO_MORE
		return
	}

	sum := 0. //sum
	for _, f := range data {
		sum += f
	}
	ave = sum / float64(n)
	ep := 0.
	for _, f := range data {
		s := f - ave

		adev += fabs(s)
		ep += s
		p := s * s
		variance += p
		p *= s
		skew += p
		p *= s
		curt += p
	}
	adev /= fn
	variance = (variance - ep*ep/fn) / (fn - 1)
	sdev = sqrt(variance)

	if variance == 0. {
		err = nerror(`No skew/kurtosis when variance = 0 (in moment)`)
		return
	}
	skew /= (fn * variance * sdev)
	curt = curt/(fn*variance*variance) - 3.0
	return
}
