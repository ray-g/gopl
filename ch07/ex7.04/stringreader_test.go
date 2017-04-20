package stringreader

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	s := "hello 世界"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Errorf("n=%d err=%s", n, err)
	}
	if b.String() != s {
		t.Errorf("\"%s\" != \"%s\"", b.String(), s)
	}
}

func TestNewReaderWithHTML(t *testing.T) {
	s := "<html><body><p>hi</p></body></html>"
	_, err := html.Parse(NewReader(s))
	if err != nil {
		t.Error()
	}
}
