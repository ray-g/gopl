package popcount

// PopCountClears returns the population count (number of set bits) of x by clearing.
func PopCountClears(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
