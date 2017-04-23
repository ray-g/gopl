package main

import (
	"net/url"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestRewriteLocalLinks(t *testing.T) {
	url_, err := url.Parse("http://www.test.com/a/b/c")
	if err != nil {
		t.Fatal(err)
	}
	base = url_
	input := `<html><body><a href="http://www.test.com/d/e">text</a></body></html>`
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	nodes := linkNodes(doc)
	rewriteLocalLinks(nodes, url_)
	want := "/d/e"
	got := nodes[0].Attr[0].Val
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
