package popcount

import "testing"

func benchmarkPopCount(b *testing.B, f func(uint64) int, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(uint64(j))
		}
	}
}

func benchmarkPopCountTable(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := range pc {
			pc[j] = pc[j/2] + byte(j&1)
		}
		benchmarkPopCount(b, PopCountTable, n)
	}
}

func BenchmarkPopCountTable1(b *testing.B)      { benchmarkPopCountTable(b, 1) }
func BenchmarkPopCountTable10(b *testing.B)     { benchmarkPopCountTable(b, 10) }
func BenchmarkPopCountTable100(b *testing.B)    { benchmarkPopCountTable(b, 100) }
func BenchmarkPopCountTable1000(b *testing.B)   { benchmarkPopCountTable(b, 1000) }
func BenchmarkPopCountTable10000(b *testing.B)  { benchmarkPopCountTable(b, 10000) }
func BenchmarkPopCountTable100000(b *testing.B) { benchmarkPopCountTable(b, 100000) }

func BenchmarkPopCountShift1(b *testing.B)      { benchmarkPopCount(b, PopCountShift, 1) }
func BenchmarkPopCountShift10(b *testing.B)     { benchmarkPopCount(b, PopCountShift, 10) }
func BenchmarkPopCountShift100(b *testing.B)    { benchmarkPopCount(b, PopCountShift, 100) }
func BenchmarkPopCountShift1000(b *testing.B)   { benchmarkPopCount(b, PopCountShift, 1000) }
func BenchmarkPopCountShift10000(b *testing.B)  { benchmarkPopCount(b, PopCountShift, 10000) }
func BenchmarkPopCountShift100000(b *testing.B) { benchmarkPopCount(b, PopCountShift, 100000) }

func BenchmarkPopCountClears1(b *testing.B)      { benchmarkPopCount(b, PopCountClears, 1) }
func BenchmarkPopCountClears10(b *testing.B)     { benchmarkPopCount(b, PopCountClears, 10) }
func BenchmarkPopCountClears100(b *testing.B)    { benchmarkPopCount(b, PopCountClears, 100) }
func BenchmarkPopCountClears1000(b *testing.B)   { benchmarkPopCount(b, PopCountClears, 1000) }
func BenchmarkPopCountClears10000(b *testing.B)  { benchmarkPopCount(b, PopCountClears, 10000) }
func BenchmarkPopCountClears100000(b *testing.B) { benchmarkPopCount(b, PopCountClears, 100000) }
