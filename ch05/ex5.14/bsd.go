package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

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

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func main() {
	for corse, prereq := range prereqs {
		fmt.Fprintf(stdout, "==========\n\"%s\" depends on following corses:\n", corse)
		breadthFirst(printDepends, prereq)
		fmt.Fprintf(stdout, "==========\n\n")
	}
}

func printDepends(corse string) (order []string) {
	fmt.Fprintf(stdout, "%s\n", corse)
	for _, item := range prereqs[corse] {
		order = append(order, item)
	}
	return
}
