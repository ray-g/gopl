package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	b := &bytes.Buffer{}
	b.WriteString(c.fn)
	b.WriteByte('(')
	for i, a := range c.args {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(a.String())
	}
	b.WriteByte(')')
	return b.String()
}
