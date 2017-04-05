// Exercise 1.2 Test echo performance
package echo

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

// EchoLen use len to loop os.Args and contruct string directly
func EchoLen() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Fprintln(out, s)
}

// EchoRange use Range to loop os.Args and contruct string directly
func EchoRange() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
}

// EchoJoin strings.Join to print os.Args
func EchoJoin() {
	fmt.Fprintln(out, strings.Join(os.Args[1:], " "))
}
