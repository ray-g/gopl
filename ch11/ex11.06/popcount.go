package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

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

// PopCountClears returns the population count (number of set bits) of x by clearing.
func PopCountClears(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
