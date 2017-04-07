# Exercise 3.8 (P63)

Rendering fractals at high zoom levels demands great arithmetic precision.
Implement the same fractal using four different representations of numbers:
`complex64`, `complex128`, `big.Float`, and `big.Rat`.
(The latter two types are found in the `math/big` package.
`Float` uses arbitrary but bounded-precision floating-point;
`Rat` uses unbounded-precision rational numbers)
How do theycompare in performance and memory usage?
At what zoom levels do rendering artifacts become visible?
