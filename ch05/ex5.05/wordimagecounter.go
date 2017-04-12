package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[len(u)-1]
		u = u[:len(u)-1]
		switch n.Type {
		case html.TextNode:
			if n.Parent.Data != "script" && n.Parent.Data != "style" {
				words += wordCount(n.Data)
			}
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
	}
	return
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(stderr, "usage: %s URL\n", os.Args[0])
		return
	}
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(stdout, "words: %d\nimages: %d\n", words, images)
}
