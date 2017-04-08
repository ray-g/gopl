package unispace

import "testing"

func TestUniSpace(t *testing.T) {
	tcs := []struct {
		input   string
		expects string
	}{
		{"Hello　世界", "Hello 世界"},
	}

	for _, tc := range tcs {
		ret := string(uniSpace([]byte(tc.input)))
		if ret != tc.expects {
			t.Errorf("Failed to remove unicode space. Input: %s, expects: %s, results: %s", tc.input, tc.expects, ret)
		}
	}
}
