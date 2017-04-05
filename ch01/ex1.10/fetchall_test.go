package main

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func cleanUp() {
	for i := 0; i < 10; i++ {
		fileName := fileBase + strconv.Itoa(i) + fileExt
		if fileExists(fileName) {
			os.Remove(fileName)
		}
	}
}

func TestFetchAll(t *testing.T) {
	var tcs = []struct {
		args    []string
		expects string
	}{
		{[]string{"fetchall", "http://github.com"}, "http://github.com"},
		{[]string{"fetchall", "http://"}, "no Host in request URL"},
		{[]string{"fetchall", "http://bad.gopl.io"}, "no such host"},
	}

	for _, tc := range tcs {
		os.Args = tc.args

		stdout = new(bytes.Buffer) // captured output
		main()
		ret := stdout.(*bytes.Buffer).String()
		if !strings.Contains(ret, tc.expects) {
			t.Errorf("Failed to fetch. Result = %q, Expects %q", ret, tc.expects)
		}
	}

	cleanUp()
}
