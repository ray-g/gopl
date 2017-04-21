# Exercise 7.10 (P191)

The `sort.Interface` type can be adapted to other use.
Write a function `IsPalidrome(s sort.Interface) bool` that reports whether the sequence `s` is a palindrome,
in other words, reversig the sequence would not change it.
Assume that the elements at indices `i` and `j` are equal if `!s.Less(i, j) && !s.Less(j, i)`
