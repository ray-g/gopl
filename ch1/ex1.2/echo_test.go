package main

import (
	"bytes"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	tcs := []struct {
		args    []string
		expects string
	}{
		{[]string{"echo"}, "[0] echo\n"},
		{[]string{"echo", "1", "2", "3"}, "[0] echo\n[1] 1\n[2] 2\n[3] 3\n"},
		{[]string{"echo", "a", "b", "c"}, "[0] echo\n[1] a\n[2] b\n[3] c\n"},
	}

	for _, tc := range tcs {
		os.Args = tc.args
		out = new(bytes.Buffer)
		main()
		ret := out.(*bytes.Buffer).String()
		if ret != tc.expects {
			t.Errorf("Echo Failed. Expects: \"%s\", Got: \"%s\"", tc.expects, ret)
		}
	}
}
