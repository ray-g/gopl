package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettify(t *testing.T) {
	input := `
<html>
<body>
	<p class="something" id="short"><span class="special">hi</span></p><br/>
</body>
</html>
`
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}
	stdout = new(bytes.Buffer)
	prettify(doc)
	_, err = html.Parse(bytes.NewReader(stdout.(*bytes.Buffer).Bytes()))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestMain(t *testing.T) {
	input := `
<html>
<body>
	<p class="something" id="short"><span class="special">hi</span></p><br/>
</body>
</html>
`
	stdin = strings.NewReader(input)
	stdout = new(bytes.Buffer)
	main()
	_, err := html.Parse(bytes.NewReader(stdout.(*bytes.Buffer).Bytes()))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
