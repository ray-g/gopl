package splitest

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tcs = []struct {
		str     string
		sep     string
		expects int
	}{
		{"a:b:c", ":", 3},
		{"a:b:c", ",", 1},
		{"a,b,c", ",", 3},
		{"a,b,c", ":", 1},
	}
	for _, tc := range tcs {
		words := strings.Split(tc.str, tc.sep)
		if actual := len(words); actual != tc.expects {
			t.Errorf("Split(%q, %q), expects: %d, actual: %d.", tc.str, tc.sep, tc.expects, actual)
		}
	}
}
