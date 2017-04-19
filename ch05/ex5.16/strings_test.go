package strings

import "testing"

func TestJoin(t *testing.T) {
	tcs := []struct {
		sep    string
		params []string
		expect string
	}{
		{"", []string{}, ""},
		{" ", []string{}, ""},
		{"+", []string{}, ""},
		{"+", []string{"aaa"}, "aaa"},
		{"+", []string{"aaa", "bbb"}, "aaa+bbb"},
		{"+", []string{"aaa", "bbb", "ccc"}, "aaa+bbb+ccc"},
		{"+", []string{"aaa", "bbb", "ccc", "ddd"}, "aaa+bbb+ccc+ddd"},
	}

	for _, tc := range tcs {
		actual := Join(tc.sep, tc.params...)
		if actual != tc.expect {
			t.Errorf("Sep: \"%s\", Params: %q, Expects: \"%s\", Actual: \"%s\"", tc.sep, tc.params, tc.expect, actual)
		}
	}
}
