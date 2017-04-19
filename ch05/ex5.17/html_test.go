package html

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByTagName(t *testing.T) {
	source, _ := ioutil.ReadFile("test.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	images := ElementsByTagName(doc, "img")
	if len(images) != 2 {
		t.Errorf("Not found images: %q", images)
	}

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	if len(headings) != 4 {
		t.Errorf("Not found headings: %q", headings)
	}

	empty := ElementsByTagName(doc)
	if len(empty) > 0 {
		t.Errorf("Found empty: %q", empty)
	}
}
