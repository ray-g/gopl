package main

import (
	"bytes"
	"os"
	"testing"
)

func TestComma(t *testing.T) {
	var tcs = []struct {
		args    []string
		expects string
	}{
		{[]string{"comma", "1", "+12", "-123", "-12.34", "+12345678.09"}, "  1\n  +12\n  -123\n  -12.34\n  +12,345,678.09\n"},
		{[]string{"comma", "1234567"}, "  1,234,567\n"},
	}

	for _, tc := range tcs {
		os.Args = tc.args
		stdout = new(bytes.Buffer) // captured output
		main()
		ret := stdout.(*bytes.Buffer).String()
		if ret != tc.expects {
			t.Errorf("Failed comma. Value: %v, Expects %q, Result = %q", tc.args, tc.expects, ret)
		}
	}
}
