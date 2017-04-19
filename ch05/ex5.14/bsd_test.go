package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestBreadthFirst(t *testing.T) {
	var tests = []struct {
		corse   string
		expects []string
	}{
		{"discrete math", []string{"intro to programming"}},
		{"programming languages", []string{"data structures", "computer organization", "discrete math", "intro to programming"}},
	}

	for _, test := range tests {
		stdout = new(bytes.Buffer)
		breadthFirst(printDepends, prereqs[test.corse])
		actual := stdout.(*bytes.Buffer).String()
		for _, expect := range test.expects {
			if !strings.Contains(actual, expect) {
				t.Errorf("Expects: %q\nActual: %q", expect, actual)
			}
		}
	}

}

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer)
	main()
	actual := stdout.(*bytes.Buffer).String()
	if len(actual) == 0 {
		t.Error()
	}
}
