package palindrome

import "sort"

type nums []int

func (s nums) Len() int           { return len(s) }
func (s nums) Less(i, j int) bool { return s[i] < s[j] }
func (s nums) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// IsPalindrome returns slice s is a palindrome or not.
func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - 1 - i
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
