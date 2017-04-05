// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout
var in io.Reader = os.Stdin
var err = os.Stderr

func main() {
	// {filename: {line: count} }
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		if _, ok := in.(*os.File); ok {
			stat, _ := in.(*os.File).Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				// fmt.Println("data is being piped to stdin")
			} else {
				// fmt.Println("stdin is from a terminal")
				fmt.Println("Ctrl-D to finish input or Ctrl-C to terminate.")
			}
		}

		countLines("*Stdin", in, counts)
	} else {
		for _, arg := range files {
			f, e := os.Open(arg)
			if e != nil {
				fmt.Fprintf(err, "dup2: %v\n", e)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	for file, dup := range counts {
		for line, n := range dup {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, file)
			}
		}
	}
}

func countLines(file string, f io.Reader, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[file]; !ok {
			counts[file] = make(map[string]int)
		}
		counts[file][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
