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
	counts := make(map[rune]int)       // counts of Unicode characters
	categories := make(map[string]int) // counts of Unicode characters categories
	var utflen [utf8.UTFMax + 1]int    // count of lengths of UTF-8 encodings
	invalid := 0                       // count of invalid UTF-8 characters

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
		countType(r, categories)
		counts[r]++
		utflen[n]++
	}
	fmt.Fprintf(stdout, "rune\tcount\n")
	// sort runes to have a better print order
	var runes []rune
	for r := range counts {
		runes = append(runes, r)
	}
	sort.Sort(RuneSlice(runes))
	for _, r := range runes {
		fmt.Fprintf(stdout, "%q\t%d\n", r, counts[r])
	}
	fmt.Fprint(stdout, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(stdout, "%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Fprintf(stdout, "\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Fprint(stdout, "\ncategory\tcount\n")
	// sort keys to have a better print order
	var keys []string
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(stdout, "%s\t%d\n", k, categories[k])
	}
}

func countType(r rune, categories map[string]int) {
	if unicode.IsControl(r) {
		categories["Control"]++
	}
	if unicode.IsDigit(r) {
		categories["Digit"]++
	}
	if unicode.IsGraphic(r) {
		categories["Graphic"]++
	}
	if unicode.IsLetter(r) {
		categories["Letter"]++
	}
	if unicode.IsLower(r) {
		categories["Lower"]++
	}
	if unicode.IsMark(r) {
		categories["Mark"]++
	}
	if unicode.IsNumber(r) {
		categories["Number"]++
	}
	if unicode.IsPrint(r) {
		categories["Printable"]++
	} else {
		categories["NonPrintable"]++
	}

	if unicode.IsPunct(r) {
		categories["Punct"]++
	}
	if unicode.IsSpace(r) {
		categories["Space"]++
	}
	if unicode.IsSymbol(r) {
		categories["Symbol"]++
	}
	if unicode.IsTitle(r) {
		categories["Title"]++
	}
	if unicode.IsUpper(r) {
		categories["Upper"]++
	}
}
