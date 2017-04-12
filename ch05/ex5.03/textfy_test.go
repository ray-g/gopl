package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestMain(t *testing.T) {
	expects := "The Go Programming Language\n...\nThe Go Programming Language\nPackages\nPlay\nBuild version go1.8.1.\n"
	source, _ := ioutil.ReadFile("./golang.org.test.html")
	stdin = bytes.NewBuffer(source)
	stdout = new(bytes.Buffer)
	main()
	ret := stdout.(*bytes.Buffer).String()
	if ret != expects {
		t.Errorf("Expects:\n%v\n\nResults:\n%v", expects, ret)
	}
}
