package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

func main() {
	h1 := sha256.Sum256([]byte("abc"))
	h2 := sha256.Sum256([]byte("ABC"))

	fmt.Fprintf(stdout, "Different bit count is %d\n", countDiffBits(h1, h2))
}

func countDiffBits(sha1, sha2 [32]uint8) int {
	n := 0
	for i := 0; i < len(sha1); i++ {
		s := sha1[i] ^ sha2[i]
		n += bitCount(s)
	}
	return n
}

func bitCount(x uint8) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
