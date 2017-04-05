package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestFetchAll(t *testing.T) {
	var tcs = []struct {
		args    []string
		expects string
	}{
		{[]string{"fetchall", "http://github.com"}, "request canceled"},
	}

	for _, tc := range tcs {
		os.Args = tc.args

		stdout = new(bytes.Buffer) // captured output
		close(done)
		main()
		ret := stdout.(*bytes.Buffer).String()
		if !strings.Contains(ret, tc.expects) {
			t.Errorf("Failed to fetch. Result = %q, Expects %q", ret, tc.expects)
		}
	}
}
