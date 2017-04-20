package counterwriter

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var tcs = []struct {
		inputs  [][]byte
		expects int64
	}{
		{[][]byte{[]byte("")}, 0},
		{[][]byte{[]byte("one"), []byte("two")}, 6},
	}
	for _, tc := range tcs {
		w, c := CountingWriter(new(bytes.Buffer))
		for _, p := range tc.inputs {
			w.Write(p)
		}
		if *c != tc.expects {
			t.Errorf("Inputs: %v, Expects: %v, Actual %v", tc.inputs, tc.expects, *c)
		}
	}
}
