package anagram

import (
	"reflect"
	"sort"
	"strings"
)

func isAnagramReflect(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	a1 := strings.Split(s1, "")
	a2 := strings.Split(s2, "")

	sort.Strings(a1)
	sort.Strings(a2)

	return reflect.DeepEqual(a1, a2)
}

func isAnagramMap(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, c := range s1 {
		m1[c]++
	}

	for _, c := range s2 {
		m2[c]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func isAnagramRemove(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	for _, c := range s1 {
		if strings.Contains(s2, string(c)) {
			s2 = strings.Replace(s2, string(c), "", 1)
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}
