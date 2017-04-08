package reverse

import "testing"

func TestReverse(t *testing.T) {
	tcs := []struct {
		input   [10]int
		expects [10]int
	}{
		{[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, [10]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tc := range tcs {
		input := tc.input
		reverse(&input)
		for i := 0; i < len(tc.input); i++ {
			if input[i] != tc.expects[i] {
				t.Errorf("Failed reverse. input: %v, expects: %v, results: %v", tc.input, tc.expects, input)
			}
		}
	}
}
