package reverse

import "testing"

func TestReverse(t *testing.T) {
	tcs := []struct {
		input   string
		expects string
	}{
		{"Hello 世界", "界世 olleH"},
	}

	for _, tc := range tcs {
		b := []byte(tc.input)
		reverse(b)
		ret := string(b)
		if ret != tc.expects {
			t.Errorf("Failed to remove unicode space. Input: %s, expects: %s, results: %s", tc.input, tc.expects, ret)
		}
	}
}
