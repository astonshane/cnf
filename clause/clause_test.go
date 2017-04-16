package clause

import (
	"cnf/element"
	"cnf/literal"
	"testing"
)

var la = literal.Literal{Name: "A"}
var lb = literal.Literal{Name: "B"}

var c1 = Clause{Symbol: "v", elements: []element.Element{la, lb}}
var c2 = Clause{Symbol: "~", elements: []element.Element{la}}
var c3 = Clause{Symbol: "v", elements: []element.Element{la, c2}}
var c4 = Clause{Symbol: "->", elements: []element.Element{la, lb}}

type testpair struct {
	symbol   string
	elements []element.Element // Append()
	str      string            // String()
	cnf      bool              // isCnf()
	nl       bool              // isNegativeLiteral()
}

var testpairs = []testpair{
	{"~", []element.Element{la}, "~[A]", true, true},
	{"~", []element.Element{c4}, "~[->[A B]]", false, false},
	{"^", []element.Element{la, lb}, "^[A B]", true, false},
	{"^", []element.Element{la, c1}, "^[A v[A B]]", true, false},
	{"^", []element.Element{c3, la}, "^[v[A ~[A]] A]", true, false},
	{"^", []element.Element{c4, la}, "^[->[A B] A]", false, false},
	{"v", []element.Element{la, lb}, "v[A B]", true, false},
	{"v", []element.Element{c4, la}, "v[->[A B] A]", false, false},
	{"->", []element.Element{la, lb}, "->[A B]", false, false},
	{"<->", []element.Element{la, lb}, "<->[A B]", false, false},
}

func TestClause(t *testing.T) {
	for _, pair := range testpairs {
		c := Clause{Symbol: pair.symbol}
		for _, elm := range pair.elements {
			c = c.Append(elm)
		}

		gotString := c.String()
		if gotString != pair.str {
			t.Error(
				"For symbol", pair.symbol,
				"and elements", pair.elements,
				"Expected", pair.str,
				"Got String()", gotString,
			)
		}

		gotIsCnf := c.IsCnf()
		if gotIsCnf != pair.cnf {
			t.Error(
				"For symbol", pair.symbol,
				"and elements", pair.elements,
				"Expected", pair.cnf,
				"Got IsCnf()", gotIsCnf,
			)
		}

		gotNegLit := c.IsNegativeLiteral()
		if gotNegLit != pair.nl {
			t.Error(
				"For symbol", pair.symbol,
				"and elements", pair.elements,
				"Expected", pair.nl,
				"Got IsNegativeLiteral()", gotNegLit,
			)
		}
	}
}
