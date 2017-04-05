package main

import (
	"bytes"
	"testing"
)

func TestLissajous(t *testing.T) {
	out = new(bytes.Buffer) // captured output
	main()
	ret := out.(*bytes.Buffer).String()
	if len(ret) == 0 {
		t.Error("Error generate GIF...")
	}
}
