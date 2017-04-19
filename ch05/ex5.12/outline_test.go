package main

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func outlineOri(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

var gDepth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(stdout, "%*s<%s>\n", gDepth*2, "", n.Data)
		gDepth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		gDepth--
		fmt.Fprintf(stdout, "%*s</%s>\n", gDepth*2, "", n.Data)
	}
}

func TestOutline(t *testing.T) {
	var tests = []struct {
		url string
	}{
		{"https://github.com/"},
	}

	for _, test := range tests {
		stdout = new(bytes.Buffer)
		outlineOri(test.url) //original implementation
		expects := stdout.(*bytes.Buffer).String()

		stdout = new(bytes.Buffer)
		err := outline(test.url)
		actual := stdout.(*bytes.Buffer).String()

		if err != nil || actual != expects {
			t.Errorf("Expects:\n%v\nActual:\n%v", expects, actual)
		}
	}
}
