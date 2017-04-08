package main

import (
	"bytes"
	"os"
	"testing"
)

func TestWordFreq(t *testing.T) {
	tcs := []struct {
		args       []string
		expectsOut string
		expectsErr string
	}{
		{[]string{"wordfreq"}, "wordfreq <filename>\n", ""},
		{[]string{"wordfreq", "no_file"}, "", "open no_file: no such file or directory"},
		{[]string{"wordfreq", "foo.txt"},
			"word\tcount\n" +
				"Hello\t1\n" +
				"World\t1\n" +
				"hello\t1\n" +
				"world\t1\n" +
				"世界\t2\n" +
				"你好\t2\n", ""},
	}

	for _, tc := range tcs {
		os.Args = tc.args
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
		main()
		retOut := stdout.(*bytes.Buffer).String()
		retErr := stderr.(*bytes.Buffer).String()

		if retOut != tc.expectsOut || retErr != tc.expectsErr {
			t.Errorf("Failed count words. Args: %v, Expect Output: %q, Expect Error: %q, Result Output: %q, Result Error: %q", tc.args, tc.expectsOut, tc.expectsErr, retOut, retErr)
		}

	}
}
