package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	expects := "Different bit count is 126\n"
	main()
	ret := stdout.(*bytes.Buffer).String()
	if ret != expects {
		t.Errorf("Results = %q, Expects %q", ret, expects)
	}
}

func TestCountDiffBits(t *testing.T) {
	var tcs = []struct {
		hash1   [32]uint8
		hash2   [32]uint8
		expects int
	}{
		{[...]uint8{31: 1}, [...]uint8{31: 0}, 1},
		{[...]uint8{31: 5}, [...]uint8{31: 1}, 1},
		{[...]uint8{0: 11, 30: 2, 31: 5}, [...]uint8{31: 1}, 5},
	}

	for _, tc := range tcs {
		ret := countDiffBits(tc.hash1, tc.hash2)
		if ret != tc.expects {
			t.Errorf("Result = %v, Expects %v", ret, tc.expects)
		}
	}
}
