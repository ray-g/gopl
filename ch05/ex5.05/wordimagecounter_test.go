package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
	source, _ := ioutil.ReadFile("./test.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	wordsExpects := 5
	imagesExpects := 2
	words, images := countWordsAndImages(doc)
	if words != wordsExpects {
		t.Errorf("Words count wrong. Expects: %d, got %d\n", wordsExpects, words)
	}

	if images != imagesExpects {
		t.Errorf("Images count wrong. Expects: %d, got %d\n", imagesExpects, images)
	}
}

func TestMain(t *testing.T) {
	os.Args = []string{"", "https://github.com/"}
	stdout = new(bytes.Buffer)
	main()
	ret := stdout.(*bytes.Buffer).String()
	if len(ret) == 0 {
		t.Errorf("No output")
	}

	os.Args = []string{"wiCounter"}
	stderr = new(bytes.Buffer)
	eexpects := "usage: wiCounter URL\n"
	main()
	eret := stderr.(*bytes.Buffer).String()
	if eret != eexpects {
		t.Error()
	}
}
