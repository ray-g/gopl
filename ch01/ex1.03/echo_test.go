package echo

import (
	"bytes"
	"os"
	"testing"
)

func makeTestCases() []struct {
	args    []string
	expects string
} {
	return []struct {
		args    []string
		expects string
	}{
		{[]string{"echo"}, "\n"},
		{[]string{"echo", "1", "2", "3"}, "1 2 3\n"},
		{[]string{"echo", "a", "b", "c"}, "a b c\n"},
	}
}

func TestEchoLen(t *testing.T) {
	tcs := makeTestCases()

	for _, tc := range tcs {
		os.Args = tc.args
		out = new(bytes.Buffer)
		EchoLen()
		ret := out.(*bytes.Buffer).String()
		if ret != tc.expects {
			t.Errorf("EchoLen Failed. Expects: \"%s\", Got: \"%s\"", tc.expects, ret)
		}
	}
}

func TestEchoRange(t *testing.T) {
	tcs := makeTestCases()

	for _, tc := range tcs {
		os.Args = tc.args
		out = new(bytes.Buffer)
		EchoRange()
		ret := out.(*bytes.Buffer).String()
		if ret != tc.expects {
			t.Errorf("EchoRange Failed. Expects: \"%s\", Got: \"%s\"", tc.expects, ret)
		}
	}
}

func TestEchoJoin(t *testing.T) {
	tcs := makeTestCases()

	for _, tc := range tcs {
		os.Args = tc.args
		out = new(bytes.Buffer)
		EchoJoin()
		ret := out.(*bytes.Buffer).String()
		if ret != tc.expects {
			t.Errorf("EchoJoin Failed. Expects: \"%s\", Got: \"%s\"", tc.expects, ret)
		}
	}
}

func makeBenchmarkStrings(size int) []string {
	strs := make([]string, size)
	for i := range strs {
		if (i & 1) == 0 {
			strs[i] = "foo"
		} else {
			strs[i] = "bar"
		}
	}
	return strs
}

func benchmarkEchoLen(b *testing.B, size int) {
	strs := makeBenchmarkStrings(size)
	os.Args = strs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoLen()
	}
}

func benchmarkEchoRange(b *testing.B, size int) {
	strs := makeBenchmarkStrings(size)
	os.Args = strs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoRange()
	}
}

func benchmarkEchoJoin(b *testing.B, size int) {
	strs := makeBenchmarkStrings(size)
	os.Args = strs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoJoin()
	}
}

func BenchmarkEchoLen10(b *testing.B)     { benchmarkEchoLen(b, 10) }
func BenchmarkEchoLen100(b *testing.B)    { benchmarkEchoLen(b, 100) }
func BenchmarkEchoLen1000(b *testing.B)   { benchmarkEchoLen(b, 1000) }
func BenchmarkEchoLen10000(b *testing.B)  { benchmarkEchoLen(b, 10000) }
func BenchmarkEchoLen100000(b *testing.B) { benchmarkEchoLen(b, 100000) }

func BenchmarkEchoRange10(b *testing.B)     { benchmarkEchoRange(b, 10) }
func BenchmarkEchoRange100(b *testing.B)    { benchmarkEchoRange(b, 100) }
func BenchmarkEchoRange1000(b *testing.B)   { benchmarkEchoRange(b, 1000) }
func BenchmarkEchoRange10000(b *testing.B)  { benchmarkEchoRange(b, 10000) }
func BenchmarkEchoRange100000(b *testing.B) { benchmarkEchoRange(b, 100000) }

func BenchmarkEchoJoin10(b *testing.B)     { benchmarkEchoJoin(b, 10) }
func BenchmarkEchoJoin100(b *testing.B)    { benchmarkEchoJoin(b, 100) }
func BenchmarkEchoJoin1000(b *testing.B)   { benchmarkEchoJoin(b, 1000) }
func BenchmarkEchoJoin10000(b *testing.B)  { benchmarkEchoJoin(b, 10000) }
func BenchmarkEchoJoin100000(b *testing.B) { benchmarkEchoJoin(b, 100000) }
