package minmax

import "testing"

func TestMin(t *testing.T) {
	tcs := []struct {
		params []int
		expect int
	}{
		{[]int{1, 2}, 1},
		{[]int{2, 4, 1, 3}, 1},
	}

	for _, tc := range tcs {
		actual := min(tc.params...)
		if actual != tc.expect {
			t.Errorf("Params: %q, Expects: %d, Actual: %d", tc.params, tc.expect, actual)
		}
	}
}

func TestMax(t *testing.T) {
	tcs := []struct {
		params []int
		expect int
	}{
		{[]int{1, 2}, 2},
		{[]int{2, 4, 1, 3}, 4},
	}

	for _, tc := range tcs {
		actual := max(tc.params...)
		if actual != tc.expect {
			t.Errorf("Params: %q, Expects: %d, Actual: %d", tc.params, tc.expect, actual)
		}
	}
}

func TestMin2(t *testing.T) {
	tcs := []struct {
		base   int
		params []int
		expect int
	}{
		{0, []int{}, 0},
		{3, []int{1, 2}, 1},
		{0, []int{2, 4, 1, 3}, 0},
	}

	for _, tc := range tcs {
		actual := min2(tc.base, tc.params...)
		if actual != tc.expect {
			t.Errorf("Base %d, Params: %q, Expects: %d, Actual: %d", tc.base, tc.params, tc.expect, actual)
		}
	}
}

func TestMax2(t *testing.T) {
	tcs := []struct {
		base   int
		params []int
		expect int
	}{
		{0, []int{}, 0},
		{0, []int{1, 2}, 2},
		{0, []int{2, 4, 1, 3}, 4},
	}

	for _, tc := range tcs {
		actual := max2(tc.base, tc.params...)
		if actual != tc.expect {
			t.Errorf("Base %d, Params: %q, Expects: %d, Actual: %d", tc.base, tc.params, tc.expect, actual)
		}
	}
}

func TestMinPanic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			if p != "No Args" {
				t.Errorf("Wrong Panic: %v", p)
			}
		} else {
			t.Error("No panic occured")
		}
	}()
	min()
}

func TestMaxPanic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			if p != "No Args" {
				t.Errorf("Wrong Panic: %v", p)
			}
		} else {
			t.Error("No panic occured")
		}
	}()
	max()
}
