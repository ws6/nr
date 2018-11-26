package nr

//Chsone chi-square implementation 1
//Given the array bins[1..nbins] containing the observed numbers of events, and an array
//ebins[1..nbins] containing the expected numbers of events, and given the number of constraints
//knstrn (normally one), this routine returns (trivially) the number of degrees of freedom
//df, and (nontrivially) the chi-square chsq and the significance prob. A small value of prob
//indicates a significant difference between the distributions bins and ebins. Note that bins
//and ebins are both float arrays, although bins will normally contain integer values.
