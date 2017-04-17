package main

import (
	"bytes"
	"testing"
)

func TestTopoSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		corses := topoSort(prereqs)
		for n, corse := range corses {
			studied := make(map[string]bool)
			for _, c := range corses[:n] {
				studied[c] = true
			}
			requests := prereqs[corse]
			for _, request := range requests {
				if _, ok := studied[request]; ok != true {
					t.Errorf("%s is requested to study %s first", corse, request)
				}

			}
		}
	}

}

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer)
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Error()
	}
}
