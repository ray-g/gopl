// Exercise 1.2 Echo to print index and value one per line
package main

import (
	"fmt"
	"io"
	"os"
)

// Create Stdout for testing
var out io.Writer = os.Stdout

//!+
func main() {
	for index, arg := range os.Args {
		fmt.Fprintf(out, "[%d] %s\n", index, arg)
	}
}
