package nr

//Returns in c[1..np], in wrap-around order (N.B.!) consistent with the argument respns in
//routine convlv, a set of Savitzky-Golay filter coefficients. nl is the number of leftward (past)
//data points used, while nr is the number of rightward (future) data points, making the total
//number of data points used nl+nr+1. ld is the order of the derivative desired (e.g., ld = 0
//for smoothed function). m is the order of the smoothing polynomial, also equal to the highest
//conserved moment; usual values are m = 2 or m = 4.
