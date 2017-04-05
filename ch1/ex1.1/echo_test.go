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
		{[]string{"echo"}, "echo\n"},
		{[]string{"echo", "1", "2", "3"}, "echo 1 2 3\n"},
		{[]string{"echo", "a", "b", "c"}, "echo a b c\n"},
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
