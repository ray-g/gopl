package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Fprintf(stdout, "  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	buf := new(bytes.Buffer)
	// Write first group of numbers
	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	// The rest
	for j := i + 3; j < n; {
		buf.WriteString("," + s[i:j])
		i, j = j, j+3
	}
	buf.WriteString("," + s[i:])
	return buf.String()
}
