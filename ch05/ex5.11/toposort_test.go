package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	stderr = new(bytes.Buffer)
	main()
	got := stderr.(*bytes.Buffer).String()
	if !strings.Contains(got, "Circle detect:") {
		t.Error()
	}
}
