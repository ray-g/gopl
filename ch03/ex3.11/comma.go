package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
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
	buf := new(bytes.Buffer)

	n := strings.Index(s, ".")
	if n == -1 {
		n = len(s)
	}

	if strings.HasPrefix(s, "+") ||
		strings.HasPrefix(s, "-") {
		buf.WriteByte(s[0])
		s = s[1:]
		n--
	}

	// Write first group of numbers
	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	// The rest
	for j, c := range s[i:n] {
		if j%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}
	buf.WriteString(s[n:])

	return buf.String()
}
