package popcount

import "testing"

func TestPopCountClears(t *testing.T) {
	tcs := []struct {
		number  uint64
		expects int
	}{
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
		{0x0000000000000002, 1},
		{0x0000000000000000, 0},
		{0x1000000000000000, 1},
	}

	for _, tc := range tcs {
		ret := PopCountClears(tc.number)
		if ret != tc.expects {
			t.Errorf("Failed PopCountClears. Number: %X, expect counts: %d, get: %d", tc.number, tc.expects, ret)
		}
	}
}

func BenchmarkPopCountClears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClears(0x1234567890ABCDEF)
	}
}
