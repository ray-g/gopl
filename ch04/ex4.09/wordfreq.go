package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(stdout, "%s <filename>\n", os.Args[0])
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(stderr, err)
		return
	}

	words := make(map[string]int)
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words[input.Text()]++
	}

	fmt.Fprint(stdout, "word\tcount\n")
	var keys []string
	for k := range words {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(stdout, "%s\t%d\n", k, words[k])
	}
}
