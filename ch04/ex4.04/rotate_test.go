package rotate

import "testing"

func TestRotate(t *testing.T) {
	tcs := []struct {
		input   []int
		order   int
		expects []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 3,
			[]int{3, 4, 5, 0, 1, 2}},

		{[]int{0, 1, 2, 3, 4, 5}, 2,
			[]int{2, 3, 4, 5, 0, 1}},

		{[]int{0, 1, 2, 3, 4, 5}, 0,
			[]int{0, 1, 2, 3, 4, 5}},

		{[]int{0, 1, 2, 3, 4, 5}, 9,
			[]int{0, 1, 2, 3, 4, 5}},

		{[]int{0, 1, 2, 3, 4, 5}, -9,
			[]int{0, 1, 2, 3, 4, 5}},
	}

	for _, tc := range tcs {
		input := make([]int, len(tc.input))
		copy(input, tc.input[:])
		rotate(input, tc.order)
		for i := 0; i < len(tc.input); i++ {
			if input[i] != tc.expects[i] {
				t.Errorf("Failed rotate. input: %v, order: %d, expects: %v, results: %v", tc.input, tc.order, tc.expects, input)
				break
			}
		}
	}
}
