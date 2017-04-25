# Exercise 9.6 (P282)

Measure how the performance of a compute-bond parallel program (see Exercise 8.5) varies with GOMAXPROCS.
What is the optimal value of your computer?
How many CPUs does your computer have?

## Result

With Intel I5 6200, 4 Cores, 2 worker threads is economical.

```text
NumCPU: 4
Done no concurrency. Used: 5.101064434s
Done. Worker Number: 1 Used: 5.076399681s
Done. Worker Number: 2 Used: 2.735404646s
Done. Worker Number: 3 Used: 2.56978686s
Done. Worker Number: 4 Used: 2.522902384s
```
