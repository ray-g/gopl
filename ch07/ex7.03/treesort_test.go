package treesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestString(t *testing.T) {
	tcs := []struct {
		values  []int
		expects string
	}{
		{[]int{1}, "[1]"},
		{[]int{1, 2, 3}, "[1 2 3]"},
		{[]int{2, 1, 3}, "[1 2 3]"},
	}

	for _, tc := range tcs {
		root := new(tree)
		for i, v := range tc.values {
			if i == 0 {
				root = &tree{value: v}
			} else {
				root = add(root, v)
			}
		}
		actual := root.String()
		if actual != tc.expects {
			t.Errorf("Tree Values: %v, Expects: %s, Actual: %s", tc.values, tc.expects, actual)
		}
	}

}
