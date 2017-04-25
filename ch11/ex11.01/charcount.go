// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr
var stdin io.Reader = os.Stdin

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	var runes []rune
	for rune := range counts {
		runes = append(runes, rune)
	}
	sort.Sort(RuneSlice(runes))
	fmt.Fprintf(stdout, "rune\tcount\n")
	for _, c := range runes {
		fmt.Fprintf(stdout, "%q\t%d\n", c, counts[c])
	}
	fmt.Fprintf(stdout, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(stdout, "%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Fprintf(stdout, "\n%d invalid UTF-8 characters\n", invalid)
	}
}
