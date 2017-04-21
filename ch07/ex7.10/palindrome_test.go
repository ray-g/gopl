package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tcs := []struct {
		s       []int
		expects bool
	}{
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 1, 1}, true},
		{[]int{1, 1}, true},
		{[]int{1}, true},
		{[]int{1, 2, 3}, false},
	}

	for _, tc := range tcs {
		actual := IsPalindrome(nums(tc.s))
		if actual != tc.expects {
			t.Errorf("%v expects: %t, actual %t", tc.s, tc.expects, actual)
		}
	}
}
