package main

import (
	"bytes"
	"testing"
)

func TestReturnByPanic(t *testing.T) {
	if r := oops(); r != 1 {
		t.Errorf("Result:%d, Expects:1", r)
	}
}

func TestMain(t *testing.T) {
	expects := "oops: 1\n"
	stdout = new(bytes.Buffer)
	main()
	actual := stdout.(*bytes.Buffer).String()
	if actual != expects {
		t.Errorf("Actual:%v, Expects:%v", actual, expects)
	}
}
