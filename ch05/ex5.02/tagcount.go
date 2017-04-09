package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"sort"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout
var stdin io.Reader = os.Stdin

func main() {
	counts, err := tagCount(stdin)
	if err != nil {
		log.Fatal(err)
	}
	var tags []string
	for tag := range counts {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	for _, tag := range tags {
		fmt.Fprintf(stdout, "%4d %s\n", counts[tag], tag)
	}
}

func tagCount(r io.Reader) (map[string]int, error) {
	counts := make(map[string]int, 0)
	t := html.NewTokenizer(r)
	var err error
	for {
		tag := t.Next()
		if tag == html.ErrorToken {
			break
		}
		name, ok := t.TagName()
		if ok {
			counts[string(name)]++
		}
	}
	if err != io.EOF {
		return counts, err
	}
	return counts, nil
}
