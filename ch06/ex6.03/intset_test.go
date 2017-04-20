package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	var tcs = []struct {
		words   []int
		expects int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 101}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 1024}, 9},
	}
	for _, tc := range tcs {
		var s IntSet
		for _, w := range tc.words {
			s.Add(w)
		}
		if s.Len() != tc.expects {
			t.Errorf("Len of %v, Actual: %d, Expects %d", tc.words, s.Len(), tc.expects)
		}
	}
}

func TestRemove(t *testing.T) {
	var tcs = []struct {
		words   []int
		remove  int
		expects string
	}{
		{[]int{}, 1, "{}"},
		{[]int{1, 2, 3}, 2, "{1 3}"},
	}

	for _, tc := range tcs {
		var s IntSet
		for _, w := range tc.words {
			s.Add(w)
		}

		s.Remove(tc.remove)

		if s.String() != tc.expects {
			t.Errorf("Remove %d from %v, Actual = %s, Expects %s", tc.remove, tc.words, s.String(), tc.expects)
		}

	}
}

func TestClear(t *testing.T) {
	var tcs = []IntSet{
		IntSet{},
		IntSet{words: []uint64{1, 2, 3}},
	}

	for _, tc := range tcs {
		tc.Clear()
		if tc.Len() != 0 {
			t.Errorf("Clear failed. IntSet: %q", tc)
		}
	}
}

func TestCopy(t *testing.T) {
	var tcs = []IntSet{
		IntSet{},
		IntSet{words: []uint64{1, 2, 3, 4, 5, 6, 7, 8}},
		IntSet{words: []uint64{1, 2, 1, 2, 1, 2, 1, 2}},
	}

	for _, tc := range tcs {
		actual := tc.Copy()
		for i := 0; i < len(tc.words); i++ {
			if actual.words[i] != tc.words[i] {
				t.Errorf("Copy: %v, Actual: %v", tc, actual)
			}
		}
	}
}

func TestAddAll(t *testing.T) {
	var tcs = []struct {
		ints    []int
		values  []int
		expects string
	}{
		{[]int{}, []int{1, 2, 3}, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{}, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{2}, "{1 2 3}"},
		{[]int{1, 2, 3, 4, 6, 7, 8}, []int{1024}, "{1 2 3 4 6 7 8 1024}"},
	}

	for _, tc := range tcs {
		s := &IntSet{}
		for _, i := range tc.ints {
			s.Add(i)
		}

		s.AddAll(tc.values...)

		if s.String() != tc.expects {
			t.Errorf("IntSet: %v AddAll %v, Expects: %s, Actual: %s", tc.ints, tc.values, tc.expects, s.String())
		}

	}
}

func TestIntersectWith(t *testing.T) {
	var tcs = []struct {
		s       []int
		t       []int
		expects string
	}{
		{[]int{1}, []int{2, 3, 4}, "{}"},
		{[]int{1, 2, 3}, []int{2, 3, 4}, "{2 3}"},
		{[]int{1, 2, 3}, []int{2}, "{2}"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 2}, "{2 8}"},
	}

	for _, tc := range tcs {
		s := &IntSet{}
		o := &IntSet{}
		for _, i := range tc.s {
			s.Add(i)
		}
		for _, i := range tc.t {
			o.Add(i)
		}

		s.IntersectWith(o)

		if s.String() != tc.expects {
			t.Errorf("%v IntersectWith %v, Expects: %s, Actual: %s", tc.s, tc.t, tc.expects, s.String())
		}

	}
}

func TestDifferenceWith(t *testing.T) {
	var tcs = []struct {
		s       []int
		t       []int
		expects string
	}{
		{[]int{1}, []int{2, 3, 4}, "{1}"},
		{[]int{1, 2, 3}, []int{2, 3, 4}, "{1}"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 2}, "{1 3 4 5 6 7}"},
	}

	for _, tc := range tcs {
		s := &IntSet{}
		o := &IntSet{}
		for _, i := range tc.s {
			s.Add(i)
		}
		for _, i := range tc.t {
			o.Add(i)
		}

		s.DifferenceWith(o)

		if s.String() != tc.expects {
			t.Errorf("%v DifferenceWith %v, Expects: %s, Actual: %s", tc.s, tc.t, tc.expects, s.String())
		}

	}
}

func TestSymmetricDifference(t *testing.T) {
	var tcs = []struct {
		s       []int
		t       []int
		expects string
	}{
		{[]int{1}, []int{2, 3, 4}, "{1 2 3 4}"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3}, []int{2}, "{1 3}"},
		{[]int{1, 2, 3}, []int{2, 3, 4}, "{1 4}"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 2, 9, 10}, "{1 3 4 5 6 7 9 10}"},
	}

	for _, tc := range tcs {
		s := &IntSet{}
		o := &IntSet{}
		for _, i := range tc.s {
			s.Add(i)
		}
		for _, i := range tc.t {
			o.Add(i)
		}

		s.SymmetricDifference(o)

		if s.String() != tc.expects {
			t.Errorf("%v SymmetricDifference with %v, Expects: %s, Actual: %s", tc.s, tc.t, tc.expects, s.String())
		}

	}
}
