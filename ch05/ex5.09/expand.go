package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr
var stdin io.Reader = os.Stdin

var regex = regexp.MustCompile(`\$\w+`)

func expand(s string, f func(string) string) string {
	wrapper := func(s string) string {
		s = s[1:]
		return f(s)
	}
	return regex.ReplaceAllStringFunc(s, wrapper)
}

func main() {
	s, err := bufio.NewReader(stdin).ReadString('.')
	if err != nil {
		fmt.Fprintf(stderr, "%s: %v", os.Args[0], err)
	}

	s = expand(s, strings.ToUpper)
	fmt.Fprintln(stdout, s)
}
