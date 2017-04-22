package main

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	xml := `<doc><a id="b"><b/>hi<b>rah</b></a></doc>`
	node, err := parse(strings.NewReader(xml))
	if err != nil {
		t.Error(err)
	}
	el := node.(*Element)
	expects := `doc []
  a [{{ id} b}]
    b []
    "hi"
    b []
      "rah"
`
	if el.String() != expects {
		t.Errorf("%q != %q", el.String(), expects)
	}
}
