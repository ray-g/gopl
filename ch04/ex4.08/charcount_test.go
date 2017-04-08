package main

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tcs = []struct {
		input   string
		expects string
	}{
		{"Hello 世界!\t你好 世界！\n",
			"rune\tcount\n" +
				"'\\t'\t1\n" +
				"'\\n'\t1\n" +
				"' '\t2\n" +
				"'!'\t1\n" +
				"'H'\t1\n" +
				"'e'\t1\n" +
				"'l'\t2\n" +
				"'o'\t1\n" +
				"'世'\t2\n" +
				"'你'\t1\n" +
				"'好'\t1\n" +
				"'界'\t2\n" +
				"'！'\t1\n" +
				"\nlen\tcount\n" +
				"1\t10\n" +
				"2\t0\n" +
				"3\t7\n" +
				"4\t0\n" +
				"\ncategory\tcount\n" +
				"Control\t2\n" +
				"Graphic\t15\n" +
				"Letter\t11\n" +
				"Lower\t4\n" +
				"NonPrintable\t2\n" +
				"Printable\t15\n" +
				"Punct\t2\n" +
				"Space\t4\n" +
				"Upper\t1\n"},
	}

	for _, tc := range tcs {
		stdin = bytes.NewBufferString(tc.input)
		stdout = new(bytes.Buffer) // captured output
		main()
		ret := stdout.(*bytes.Buffer).String()

		if ret != tc.expects {
			t.Errorf("Failed count chars.\nInput: %q,\nExpects:\n%q,\nResults:\n%q", tc.input, tc.expects, ret)
		}
	}
}
