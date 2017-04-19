package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

func oops() (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = 1
		}
	}()
	panic("oops!")
}

func main() {
	fmt.Fprintf(stdout, "oops: %d\n", oops())
}
