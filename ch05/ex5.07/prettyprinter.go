package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout
var stdin io.Reader = os.Stdin

var depth int = 0

func main() {
	doc, err := html.Parse(stdin)
	if err != nil {
		log.Fatal(err)
	}
	prettify(doc)
}

func prettify(n *html.Node) {
	start(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		prettify(c)
	}
	end(n)
}

func startElement(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}

	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	name := n.Data

	printf("%*s<%s%s%s\n", depth*2, "", name, attrStr, end)
	depth++
}

func endElement(n *html.Node) {
	depth--
	if n.FirstChild == nil {
		return
	}
	printf("%*s</%s>\n", depth*2, "", n.Data)
}

func startText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	printf("%*s%s\n", depth*2, "", n.Data)
}

func startComment(n *html.Node) {
	printf("<!--%s-->\n", n.Data)
}

func start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		startElement(n)
	case html.TextNode:
		startText(n)
	case html.CommentNode:
		startComment(n)
	}
}

func end(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		endElement(n)
	}
}

func printf(format string, args ...interface{}) {
	fmt.Fprintf(stdout, format, args...)
}
