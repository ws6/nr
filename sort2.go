package nr

//Sort2
//Sorts an array arr[1..n] into ascending order using Quicksort, while making the corresponding
//rearrangement of the array brr[1..n].
func Sort2(arr, brr []float64) error {

	if len(arr) != len(brr) {
		return nerror(`arr is not same size as brr`)
	}
	indx := Indexx(arr)

	Sort(arr)
	_brr := make([]float64, len(arr))
	for i, idx := range indx {
		_brr[i] = brr[idx]
	}

	for i, v := range _brr {
		brr[i] = v
	}

	return nil
}
