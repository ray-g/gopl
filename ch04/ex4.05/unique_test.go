package unique

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	tcs := []struct {
		input   []string
		expects []string
	}{
		{[]string{"aaa", "bbb", "bbb", "ccc", "ccc", "ccc", "ddd", "ddd", "eee"},
			[]string{"aaa", "bbb", "ccc", "ddd", "eee"}},
		{[]string{"aaa", "bbb", "bbb", "cccc", "ccc", "cccc", "ddd", "ddd", "eee"},
			[]string{"aaa", "bbb", "cccc", "ccc", "cccc", "ddd", "eee"}},
	}
	for _, tc := range tcs {
		input := make([]string, len(tc.input))
		copy(input, tc.input)
		ret := unique(input)
		if !reflect.DeepEqual(ret, tc.expects) {
			t.Errorf("Failed unique. input: %v, expects: %v, results: %v", tc.input, tc.expects, ret)
		}
	}
}
