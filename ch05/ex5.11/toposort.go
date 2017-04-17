// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":           {"data structures"},
	"calculus":             {"linear algebra"},
	"linear algebra":       {"calculus"},        // circle
	"intro to programming": {"data structures"}, // another circle

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	sortReqs, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(stderr, "%v", err)
	}
	for i, course := range sortReqs {
		fmt.Fprintf(stdout, "%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for key := range m {
		path, circle := detectCircle(key, nil, m)
		if circle {
			return nil, fmt.Errorf("Circle detect: %s", strings.Join(path, " => "))
		}
		visitAll([]string{key})
	}

	// sort.Strings(keys)
	// visitAll(keys)
	return order, nil
}

func detectCircle(key string, path []string, m map[string][]string) ([]string, bool) {
	if path == nil {
		path = append(path, key)
	}

	for _, item := range m[key] {
		for i := 0; i < len(path); i++ {
			if path[i] == item {
				path = append(path, item+" @#"+strconv.Itoa(i))
				return path, true
			}
		}

		path = append(path, item)
		return detectCircle(item, path, m)
	}

	return nil, false
}
