package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	var tcs = []struct {
		args    []string
		expects string
	}{
		{[]string{"fetch", "github.com"}, "resp.Status 200 OK\n"},
		{[]string{"fetch", "https://github.com"}, "resp.Status 200 OK\n"},
		{[]string{"fetch", "http://github.com"}, "resp.Status 200 OK\n"},
	}

	for _, tc := range tcs {
		os.Args = tc.args
		stdout = new(bytes.Buffer) // captured output
		stderr = stdout
		main()
		ret := stdout.(*bytes.Buffer).String()
		if !strings.Contains(ret, tc.expects) {
			t.Errorf("Failed Fetch url=%s, got = %q", tc.args[1], ret)
		}
	}
}
