package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestDup2(t *testing.T) {
	var tcs = []struct {
		args    []string
		input   string
		expects string
	}{
		{[]string{"dup2", "test_dup.txt", "test_nodup.txt"}, "", "ccc\t2\n\n	test_dup.txt\n"},
	}

	for _, tc := range tcs {
		os.Args = tc.args
		in = bytes.NewBufferString(tc.input)
		out = new(bytes.Buffer) // captured output
		main()
		ret := out.(*bytes.Buffer).String()
		if strings.Contains(ret, tc.expects) {
			t.Errorf("Failed Dup2. Expects = %q, Got %q", tc.expects, ret)
		}
	}
}
