package anagram

import (
	"reflect"
	"runtime"
	"testing"
)

func testIsAnagram(t *testing.T, f func(string, string) bool) {
	tcs := []struct {
		s1      string
		s2      string
		expects bool
	}{
		{"abc", "cba", true},
		{"abc", "abc", false},
		{"abc", "abcd", false},
		{"abc", "ab", false},
		{"abc", "", false},
	}

	for _, tc := range tcs {
		ret := f(tc.s1, tc.s2)
		if ret != tc.expects {
			t.Errorf("Failed %v, s1: %s, s2: %s, result: %v", getFunctionName(f), tc.s1, tc.s2, ret)
		}
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestIsAnagramReflect(t *testing.T) {
	testIsAnagram(t, isAnagramReflect)
}

func TestIsAnagramMap(t *testing.T) {
	testIsAnagram(t, isAnagramMap)
}

func TestIsAnagramRemove(t *testing.T) {
	testIsAnagram(t, isAnagramRemove)
}

func benchmarkIsAnagram(b *testing.B, f func(string, string) bool) {
	for i := 0; i < b.N; i++ {
		f("abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba")
	}
}

func BenchmarkIsAnagramReflect(b *testing.B) {
	benchmarkIsAnagram(b, isAnagramReflect)
}

func BenchmarkIsAnagramMap(b *testing.B) {
	benchmarkIsAnagram(b, isAnagramMap)
}

func BenchmarkIsAnagramRemove(b *testing.B) {
	benchmarkIsAnagram(b, isAnagramRemove)
}
