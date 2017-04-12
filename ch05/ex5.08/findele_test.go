package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	source, _ := ioutil.ReadFile("./test.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	n := ElementByID(doc, "someid")
	for _, a := range n.Attr {
		if a.Key == "href" && a.Val != "http://first" {
			t.Errorf("Wrong element: %s", a.Val)
		}
	}

}

func TestMain(t *testing.T) {
	tcs := []struct {
		args       []string
		outExpects string
		errExpects string
	}{
		{[]string{"findele"}, "", "usage: findele HTML_FILE ID\n"},
		{[]string{"findele", "./test.html"}, "", "usage: findele HTML_FILE ID\n"},
		{[]string{"findele", "./test.html", "someid"}, "ID someid found in ./test.html\n<a> has 'href' element, value is 'http://first'\n<a> has 'id' element, value is 'someid'\n", ""},
		{[]string{"findele", "./test.html", "anotherid"}, "ID anotherid found in ./test.html\n<a> has 'href' element, value is 'http://third'\n<a> has 'id' element, value is 'anotherid'\n", ""},
		{[]string{"findele", "./test.html", "notid"}, "ID notid not found in ./test.html\n", ""},
	}

	for _, tc := range tcs {
		os.Args = tc.args
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
		main()
		outGot := stdout.(*bytes.Buffer).String()
		errGot := stderr.(*bytes.Buffer).String()

		if outGot != tc.outExpects {
			t.Errorf("Output not same.\nArgs:%v\nExpects:%s\nResults:%s\n", tc.args, tc.outExpects, outGot)
		}

		if errGot != tc.errExpects {
			t.Errorf("Error not same.\nArgs:%v\nExpects:%s\nResults:%s\n", tc.args, tc.errExpects, errGot)
		}
	}
}
