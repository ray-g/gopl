package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFetch(t *testing.T) {
	var tcs = []struct {
		args []string
	}{
		{[]string{"fetch", "github.com"}},
		{[]string{"fetch", "https://github.com"}},
		{[]string{"fetch", "http://github.com"}},
	}

	for _, tc := range tcs {
		os.Args = tc.args
		stdout = new(bytes.Buffer) // captured output
		stderr = stdout
		main()
		ret := stdout.(*bytes.Buffer).String()
		if len(ret) == 0 {
			t.Errorf("Failed Fetch url=%s, got = %q", tc.args[1], ret)
		}
	}
}
