package literal

import "testing"

func TestLiteral(t *testing.T) {
	var names = []string{
		"A",
		"B",
		"CDEF",
	}
	for _, name := range names {
		l := Literal{Name: name}

		got := l.String()
		if got != name {
			t.Error("Expected", name, "Got", got)
		}

		if !l.IsCnf() {
			t.Error("Literal always should be in CNF")
		}
	}
}
