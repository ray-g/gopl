package main

import "testing"
import "reflect"

func BenchmarkUseSortByColumns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useSortByColumns()
	}
}

func BenchmarkUseSortStable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useSortStable()
	}
}

func TestSortByColumns(t *testing.T) {
	tc := useSortByColumns()
	ts := useSortStable()

	for i := 0; i < len(tc); i++ {
		if !reflect.DeepEqual(tc[i], ts[i]) {
			t.Errorf("#%d, %v != %v", i, tc[i], ts[i])
		}
	}
}
