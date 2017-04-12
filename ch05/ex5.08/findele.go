package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func ElementByID(n *html.Node, id string) *html.Node {
	if n == nil {
		return nil
	}
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}

		}
		return true
	}
	return forEachElement(n, pre, nil)
}

func forEachElement(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	nodes := make([]*html.Node, 0)
	nodes = append(nodes, n)
	for len(nodes) > 0 {
		n = nodes[0]
		nodes = nodes[1:]
		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			nodes = append(nodes, c)
		}
		if post != nil {
			if !post(n) {
				return n
			}
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(stderr, "usage: %s HTML_FILE ID\n", os.Args[0])
		return
	}
	filename := os.Args[1]
	id := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	n := ElementByID(doc, id)
	if n == nil {
		fmt.Fprintf(stdout, "ID %s not found in %s\n", id, filename)
	} else {
		fmt.Fprintf(stdout, "ID %s found in %s\n", id, filename)
		for _, a := range n.Attr {
			fmt.Fprintf(stdout, "<%s> has '%s' element, value is '%s'\n",
				n.Data, a.Key, a.Val)
		}
	}
}
