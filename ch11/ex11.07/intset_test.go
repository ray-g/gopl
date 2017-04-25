package main

import (
	"math/rand"
	"testing"
)

func newIntSets() []IntSet {
	return []IntSet{&BitIntSet{}, NewMapIntSet()}
}

func TestLenZeroInitially(t *testing.T) {
	for _, s := range newIntSets() {
		if s.Len() != 0 {
			t.Errorf("%T.Len(): got %d, want 0", s, s.Len())
		}
	}
}

func TestLenAfterAddingElements(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(2000)
		if s.Len() != 2 {
			t.Errorf("%T.Len(): got %d, want 2", s, s.Len())
		}
	}
}

func TestRemove(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Remove(0)
		if s.Has(0) {
			t.Errorf("%T: want zero removed, got %s", s, s)
		}
	}
}

func TestClear(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(1000)
		s.Clear()
		if s.Has(0) || s.Has(1000) {
			t.Errorf("%T: want empty set, got %s", s, s)
		}
	}
}

func TestCopy(t *testing.T) {
	for _, orig := range newIntSets() {
		orig.Add(1)
		copy := orig.Copy()
		copy.Add(2)
		if !copy.Has(1) || orig.Has(2) {
			t.Errorf("%T: want %s, got %s", orig, orig, copy)
		}
	}
}

func TestAddAll(t *testing.T) {
	for _, s := range newIntSets() {
		s.AddAll(0, 2, 4)
		if !s.Has(0) || !s.Has(2) || !s.Has(4) {
			t.Errorf("%T: want {2 4}, got %s", s, s)
		}
	}
}

const max = 32000

func addRandom(set IntSet, n int) {
	for i := 0; i < n; i++ {
		set.Add(rand.Intn(max))
	}
}

func benchHas(b *testing.B, set IntSet, n int) {
	addRandom(set, n)
	for i := 0; i < b.N; i++ {
		set.Has(rand.Intn(max))
	}
}

func benchAdd(b *testing.B, set IntSet, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(max))
		}
		set.Clear()
	}
}

func randInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(max)
	}
	return ints
}

func benchAddAll(b *testing.B, set IntSet, batchSize int) {
	ints := randInts(batchSize)
	for i := 0; i < b.N; i++ {
		set.AddAll(ints...)
		set.Clear()
	}
}

func benchUnionWith(bm *testing.B, a, b IntSet, n int) {
	addRandom(a, n)
	addRandom(b, n)
	for i := 0; i < bm.N; i++ {
		a.UnionWith(b)
	}
}

func benchString(b *testing.B, set IntSet, n int) {
	addRandom(set, n)
	for i := 0; i < b.N; i++ {
		set.String()
	}
}

//func Benchmark<Type><Method><Size>(b *testing.B) {
//	bench<Method>(b, New<Type>(), <Size>)
//}
func BenchmarkMapIntSetAdd10(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 10)
}
func BenchmarkMapIntSetAdd100(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 100)
}
func BenchmarkMapIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 1000)
}
func BenchmarkMapIntSetHas10(b *testing.B) {
	benchHas(b, NewMapIntSet(), 10)
}
func BenchmarkMapIntSetHas100(b *testing.B) {
	benchHas(b, NewMapIntSet(), 100)
}
func BenchmarkMapIntSetHas1000(b *testing.B) {
	benchHas(b, NewMapIntSet(), 1000)
}
func BenchmarkMapIntSetAddAll10(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 10)
}
func BenchmarkMapIntSetAddAll100(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 100)
}
func BenchmarkMapIntSetAddAll1000(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 1000)
}
func BenchmarkMapIntSetString10(b *testing.B) {
	benchString(b, NewMapIntSet(), 10)
}
func BenchmarkMapIntSetString100(b *testing.B) {
	benchString(b, NewMapIntSet(), 100)
}
func BenchmarkMapIntSetString1000(b *testing.B) {
	benchString(b, NewMapIntSet(), 1000)
}
func BenchmarkBitIntSetAdd10(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 10)
}
func BenchmarkBitIntSetAdd100(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 100)
}
func BenchmarkBitIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 1000)
}
func BenchmarkBitIntSetHas10(b *testing.B) {
	benchHas(b, NewBitIntSet(), 10)
}
func BenchmarkBitIntSetHas100(b *testing.B) {
	benchHas(b, NewBitIntSet(), 100)
}
func BenchmarkBitIntSetHas1000(b *testing.B) {
	benchHas(b, NewBitIntSet(), 1000)
}
func BenchmarkBitIntSetAddAll10(b *testing.B) {
	benchAddAll(b, NewBitIntSet(), 10)
}
func BenchmarkBitIntSetAddAll100(b *testing.B) {
	benchAddAll(b, NewBitIntSet(), 100)
}
func BenchmarkBitIntSetAddAll1000(b *testing.B) {
	benchAddAll(b, NewBitIntSet(), 1000)
}
func BenchmarkBitIntSetString10(b *testing.B) {
	benchString(b, NewBitIntSet(), 10)
}
func BenchmarkBitIntSetString100(b *testing.B) {
	benchString(b, NewBitIntSet(), 100)
}
func BenchmarkBitIntSetString1000(b *testing.B) {
	benchString(b, NewBitIntSet(), 1000)
}

func BenchMarkMapIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 10)
}
func BenchMarkMapIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 100)
}
func BenchMarkMapIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 1000)
}
func BenchMarkBitIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 10)
}
func BenchMarkBitIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 100)
}
func BenchMarkBitIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 1000)
}
