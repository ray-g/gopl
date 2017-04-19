// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var depth int

	//!+call
	forEachNode(doc,
		func(n *html.Node) { // startElement
			if n.Type == html.ElementNode {
				fmt.Fprintf(stdout, "%*s<%s>\n", depth*2, "", n.Data)
				depth++
			}
		},
		func(n *html.Node) { // endElement
			if n.Type == html.ElementNode {
				depth--
				fmt.Fprintf(stdout, "%*s</%s>\n", depth*2, "", n.Data)
			}
		})
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode
