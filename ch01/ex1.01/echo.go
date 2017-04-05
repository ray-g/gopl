// Exercise 1.1 Echo to print command name itself
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Create Stdout for testing
var out io.Writer = os.Stdout

//!+
func main() {
	fmt.Fprintln(out, strings.Join(os.Args[0:], " "))
}
