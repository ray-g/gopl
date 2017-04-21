package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tcs := []struct {
		expr   string
		env    Env
		result string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
	}
	var prevExpr string
	for _, tc := range tcs {
		// Print expr only when it changes.
		if tc.expr != prevExpr {
			fmt.Printf("\n%s\n", tc.expr)
			prevExpr = tc.expr
		}
		expr, err := Parse(tc.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		s := expr.String()
		// parsed again,
		fmt.Println("Expr!!!!!!:", s)
		reexpr, err := Parse(s)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		// yield an equivalent tree.
		got := fmt.Sprintf("%.6g", expr.Eval(tc.env))
		regot := fmt.Sprintf("%.6g", reexpr.Eval(tc.env))

		//		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != tc.result || regot != tc.result {
			t.Errorf("\n%s.Eval() in %v = %q,\n%s.Eval() in %v = %q, result %q\n",
				tc.expr, tc.env, got, reexpr, tc.env, regot, tc.result)
		}
	}
}
