# Exercise 6.5 (P168)

The type of each word used by `IntSet` is `uint64`,
but 64-bit arithmetic may be inefficient on a 32-bit platform.
Modify the program to use the `uint` type,
which is the most efficient unsigned integer type for the platform.
Instead of dividing by 64, define a constant holding the effective `size of uint` in bits, 32 or 64.
You can use the perhaps too-clever expression `32 << (^uint(0) >> 63) for this purpose.
