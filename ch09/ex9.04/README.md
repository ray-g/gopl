# Exercise 9.4 (P280)

Construct a pipeline that connects an arbitrary number of goroutines with channels.
What is the maximum number of pipeline stages you can create with out running out of memory?
How long does a value take to transit the entire pipeline?

## Result

### With ResetTimer after initialize all channels

BenchmarkPipeline10-4       500000  2573       ns/op  0   B/op  0  allocs/op
BenchmarkPipeline100-4      50000   24956      ns/op  0   B/op  0  allocs/op
BenchmarkPipeline1000-4     5000    286223     ns/op  0   B/op  0  allocs/op
BenchmarkPipeline10000-4    300     3980524    ns/op  0   B/op  0  allocs/op
BenchmarkPipeline100000-4   30      39138676   ns/op  26  B/op  0  allocs/op
BenchmarkPipeline1000000-4  3       379899748  ns/op  0   B/op  0  allocs/op

### Without ResetTimer after initialize all channels

BenchmarkPipeline10-4       500000  2497        ns/op  0          B/op  0        allocs/op
BenchmarkPipeline100-4      50000   24588       ns/op  0          B/op  0        allocs/op
BenchmarkPipeline1000-4     5000    287139      ns/op  26         B/op  0        allocs/op
BenchmarkPipeline10000-4    300     3697650     ns/op  4825       B/op  41       allocs/op
BenchmarkPipeline100000-4   30      42878492    ns/op  381184     B/op  3657     allocs/op
BenchmarkPipeline1000000-4  1       4149927880  ns/op  621803856  B/op  2799323  allocs/op

## With a 16GB OS, ~8GB available mem

When trying to create 5000000, it running out of memory.
