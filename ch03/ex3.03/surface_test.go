package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestCorner(t *testing.T) {
	var tcs = []struct {
		i       int
		j       int
		expects error
	}{
		{50, 50, errors.New("f(0, 0) was NaN. corner(50, 50)")},
		{50, 49, nil},
		{49, 50, nil},
	}

	for _, tc := range tcs {
		_, _, err := corner(tc.i, tc.j)

		if err == nil {
			if err != tc.expects {
				t.Errorf("i = %d, j = %d, expect = %v, results: %v", tc.i, tc.j, tc.expects, err)
			}
		} else {
			if err.Error() != tc.expects.Error() {
				t.Errorf("i = %d, j = %d, expect = %v, results: %v", tc.i, tc.j, tc.expects, err)
			}
		}

	}
}

func TestSurface(t *testing.T) {
	buffer := new(bytes.Buffer)
	surface(buffer)
	if strings.Contains(buffer.String(), "NaN") {
		t.Errorf("SVG included NaN value.\n%s", buffer)
	}
}
