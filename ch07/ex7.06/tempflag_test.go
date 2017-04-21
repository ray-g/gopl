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
		{[]string{"testflag", "-temp", "0K"}, "-273.15°C\n"},
		{[]string{"testflag", "-temp", "0°K"}, "-273.15°C\n"},
		{[]string{"testflag", "-temp", "0C"}, "0.00°C\n"},
		{[]string{"testflag", "-temp", "0°C"}, "0.00°C\n"},
		{[]string{"testflag", "-temp", "32F"}, "0.00°C\n"},
		{[]string{"testflag", "-temp", "32°F"}, "0.00°C\n"},
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
