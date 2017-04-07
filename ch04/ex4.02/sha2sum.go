package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

var bitlen = flag.Int("len", 256, "hash width (384 or 512)")

func main() {
	flag.Parse()
	sha2sum := func(b []byte) []byte {
		h := sha256.Sum256(b)
		return h[:]
	}

	switch *bitlen {
	case 384:
		sha2sum = func(b []byte) []byte {
			h := sha512.Sum384(b)
			return h[:]
		}
	case 512:
		fmt.Println("512")
		sha2sum = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	}

	for _, v := range flag.Args() {
		fmt.Fprintf(stdout, "%x\n", sha2sum([]byte(v)))
	}

}
