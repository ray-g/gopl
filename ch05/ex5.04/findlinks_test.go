package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func visitLoop(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}

		if n.Data == "img" || n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitLoop(links, c)
	}
	return links
}

func TestFindLinks(t *testing.T) {
	source, _ := ioutil.ReadFile("./golang.org.index.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	ret := visit(nil, doc)
	expects := visitLoop(nil, doc)

	if !reflect.DeepEqual(ret, expects) {
		t.Errorf("Results:%q\nExpects:%q", ret, expects)
	}
}

func TestMain(t *testing.T) {
	source, _ := ioutil.ReadFile("./golang.org.index.html")
	stdin = bytes.NewBuffer(source)
	stdout = new(bytes.Buffer)
	main()
	ret := stdout.(*bytes.Buffer).String()
	if len(ret) == 0 {
		t.Error()
	}
}
