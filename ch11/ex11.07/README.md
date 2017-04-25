# Exercise 11.7 (P323)

Write benchmarks for `Add`, `UnionWith`, and other methods of `*IntSet` (6.5) using large pseudo-random inputs.
How fast can you make these methods run?
How does the choice of word size affect performance?
How fast is `IntSet` compared to a set implementation based on the built-in map type?

```text
BenchmarkMapIntSetAdd10-4       1000000   1411    ns/op  323    B/op  3     allocs/op
BenchmarkMapIntSetAdd100-4      100000    16217   ns/op  3474   B/op  20    allocs/op
BenchmarkMapIntSetAdd1000-4     10000     194432  ns/op  55611  B/op  98    allocs/op
BenchmarkMapIntSetHas10-4       20000000  59.8    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetHas100-4      20000000  68.4    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetHas1000-4     20000000  63.7    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetAddAll10-4    2000000   985     ns/op  323    B/op  3     allocs/op
BenchmarkMapIntSetAddAll100-4   100000    13628   ns/op  3475   B/op  20    allocs/op
BenchmarkMapIntSetAddAll1000-4  5000      206743  ns/op  55627  B/op  98    allocs/op
BenchmarkMapIntSetString10-4    500000    2916    ns/op  368    B/op  14    allocs/op
BenchmarkMapIntSetString100-4   50000     33491   ns/op  4577   B/op  108   allocs/op
BenchmarkMapIntSetString1000-4  5000      336538  ns/op  41164  B/op  994   allocs/op
BenchmarkBitIntSetAdd10-4       3000000   536     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAdd100-4      300000    4428    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAdd1000-4     30000     45583   ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas10-4       30000000  42.8    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas100-4      30000000  41.3    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas1000-4     30000000  41.9    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll10-4    20000000  106     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll100-4   2000000   609     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll1000-4  300000    5464    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetString10-4    500000    3065    ns/op  256    B/op  12    allocs/op
BenchmarkBitIntSetString100-4   50000     27027   ns/op  3649   B/op  106   allocs/op
BenchmarkBitIntSetString1000-4  5000      209314  ns/op  33009  B/op  1002  allocs/op
```
