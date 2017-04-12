package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	tcs := []struct {
		in      string
		expects string
		foo     func(string) string
	}{
		{"following words $should $go $upper.", "following words SHOULD GO UPPER.", strings.ToUpper},
		{"FOLLOWING WORDS $SHOULD $GO $LOWER.", "FOLLOWING WORDS should go lower.", strings.ToLower},
	}

	for _, tc := range tcs {
		got := expand(tc.in, tc.foo)
		if got != tc.expects {
			t.Errorf("Failed to expand.\nInput:%s\nExpects:%s\nResults:%s\n", tc.in, tc.expects, got)
		}
	}
}

func TestMain(t *testing.T) {
	tcs := []struct {
		in      string
		expects string
		foo     func(string) string
	}{
		{"following words $should $go $upper.", "following words SHOULD GO UPPER.\n", strings.ToUpper},
	}

	for _, tc := range tcs {
		stdin = bytes.NewBufferString(tc.in)
		stdout = new(bytes.Buffer)
		main()
		got := stdout.(*bytes.Buffer).String()
		if got != tc.expects {
			t.Errorf("Failed to expand.\nInput:%s\nExpects:%s\nResults:%s\n", tc.in, tc.expects, got)
		}
	}
}
