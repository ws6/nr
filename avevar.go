package nr

func Avevar(data []float64) (
	ave float64, //average
	_var float64, //variance
	err error,
) {
	fn := float64(len(data))
	if fn <= 1. {
		err = ERR_TWO_MORE
		return
	}
	for _, f := range data {
		ave += f
	}
	ave /= fn

	ep := 0.
	s := 0.
	for _, f := range data {
		s = f - ave
		ep += s
		_var += s * s
	}

	_var = (_var - ep*ep/fn) / (fn - 1)

	return

}
