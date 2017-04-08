package main

import (
	"bytes"
	"os"
	"testing"
)

func TestIssues(t *testing.T) {
	os.Args = []string{"issues", "repo:golang/go"}
	stdout = new(bytes.Buffer)
	main()
	if len(stdout.(*bytes.Buffer).String()) <= 0 {
		t.Error()
	}
}
