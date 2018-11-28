package nr

//Crank
//Given a sorted array w[0..n-1], replaces the elements by their rank, including midranking of ties,
//and returns as s the sum of f3 − f, w here f is the number of elements in each tie.
func Crank(w []float64) (
	rank []float64, //the ranks
	s float64, //the sum of f3 − f, w here f is the number of elements in each tie.
) {
	n := len(w)
	s = 0.

	rank = make([]float64, n)
	if n == 0 {
		return
	}
	lastTieIdx := 0
	currentTieIdx := lastTieIdx
	fillRank := func() {
		_rank := .5 * float64(lastTieIdx+currentTieIdx)
		for _i := lastTieIdx; _i <= currentTieIdx; _i++ {
			rank[_i] = _rank
		}
		t := float64(currentTieIdx - lastTieIdx)

		s += t*t*t - t

	}
	for j := 0; j < n; j++ {

		if j == n-1 {
			fillRank()
			break
		}

		if w[lastTieIdx] == w[j+1] {
			currentTieIdx++

			continue

		}
		fillRank()

		//calculate the ties
		//reset lastTieIdx and lastTieIdx
		currentTieIdx = j + 1
		lastTieIdx = currentTieIdx
	}

	return
}
