# Exercise 11.6 (P323)

Write benchmarks to compare the `PopCount` implementation in Section 2.6.2 with your solutions to Exercise 2.4 and Exercise 2.5.
At what point does the table-based approach break even?

```text
BenchmarkPopCountTable1-4        20000       121925    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable10-4       10000       567430    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable100-4      10000       5678211   ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable1000-4     3000        16778123  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable10000-4    300         17531809  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable100000-4   100         54971520  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift1-4        20000000    59.1      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift10-4       3000000     620       ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift100-4      200000      8679      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift1000-4     10000       113931    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift10000-4    2000        1130659   ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift100000-4   100         11198696  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears1-4       1000000000  3.12      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears10-4      30000000    39.0      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears100-4     3000000     528       ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears1000-4    200000      7529      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears10000-4   10000       100799    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears100000-4  1000        1257457   ns/op  0  B/op  0  allocs/op
```
