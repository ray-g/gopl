package popcount

// PopCountShift returns the population count (number of set bits) of x by shifting.
func PopCountShift(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&1 != 0 {
			n++
		}
		x = x >> 1
	}
	return n
}
