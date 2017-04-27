# Exercise 12.3 (P341)

Implement the missing cases of the `encode` function.
Encode booleans at `t` and `nil`, floating-point numbers using Go's notation, and complex numbers like `1+2i` as `#C(1.0 2.0)`.
Interfaces can be encoded as a pair of a type name and a value, for instance `("[]int" (1 2 3))`, but beware that this notation is ambiguous:
the `reflect.Type.String` method may return the same string for different types.
