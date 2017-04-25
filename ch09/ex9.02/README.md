# Exercise 9.2 (P271)

Rewrite the `PopCount` example from Section 2.6.2 so that it initializes the lookup table using `sync.Once` the first time it is needed.
(Realistically, the cost of synchronization whould be prohibitive for a small and highly optimized function like `PopCount`)
