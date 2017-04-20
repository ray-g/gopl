package limitreader

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "hello 世界"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 5)
	n, _ := b.ReadFrom(r)
	if n != 5 {
		t.Errorf("n=%d", n)
	}
	if b.String() != "hello" {
		t.Errorf(`"%s" != "%s"`, b.String(), s[0:5])
	}
}
