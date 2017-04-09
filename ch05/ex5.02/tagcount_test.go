package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestTagCount(t *testing.T) {
	source, _ := ioutil.ReadFile("./golang.org.index.html")
	stdin = bytes.NewBuffer(source)
	stdout = new(bytes.Buffer)
	expects := `  22 a
  33 div
   1 form
   1 iframe
   1 input
   3 link
   3 meta
   8 option
   9 script
   3 span
   2 textarea
`
	main()
	if stdout.(*bytes.Buffer).String() != expects {
		t.Errorf("Expects:\n%q\nResults:\n%q", expects, stdout)
	}
}
