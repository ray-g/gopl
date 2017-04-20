// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset
const UINTSIZE = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/UINTSIZE, uint(x%UINTSIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/UINTSIZE, uint(x%UINTSIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll addes all the non-negative values to the set.
func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifference sets s to the symmetirc difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINTSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", UINTSIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

func bitCount(x uint) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

// Len return the number of elements
func (s *IntSet) Len() (l int) {
	for _, word := range s.words {
		l += bitCount(word)
	}
	return
}

// Remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/UINTSIZE, uint(x%UINTSIZE)
	if len(s.words) == 0 || word > len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	c := new(IntSet)
	c.words = make([]uint, len(s.words))
	copy(c.words, s.words)
	return c
}

// Elems returns a slice containing the elements of the set
func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < UINTSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, UINTSIZE*i+j)
			}
		}
	}
	return
}
