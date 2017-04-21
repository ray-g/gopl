package main

import (
	"bytes"
	"os"
	"testing"
)

func TestTempFlag(t *testing.T) {
	var tcs = []struct {
		args    []string
		expects string
	}{
		{[]string{"testflag"}, "20.00°C, Because this method is called\n"},
	}
	for _, tc := range tcs {
		os.Args = tc.args
		stdout = new(bytes.Buffer)
		main()
		actual := stdout.(*bytes.Buffer).String()
		if actual != tc.expects {
			t.Errorf("Args: %v, Expects: %v, Actual: %v", tc.args, tc.expects, actual)
		}

	}
}
