package clause

import (
	"cnf/element"
	"cnf/literal"
	"fmt"
)

// Clause type
type Clause struct {
	elements []element.Element
	Symbol   string
}

// String method
func (c Clause) String() string {
	return fmt.Sprintf("%s%v", c.Symbol, c.elements)
}

// Append to elements
func (c Clause) Append(elm element.Element) Clause {
	c.elements = append(c.elements, elm)
	return c
}

// IsCnf method
func (c Clause) IsCnf() bool {
	switch c.Symbol {
	case "~":
		return c.IsNegativeLiteral()
	case "^":
		for _, elm := range c.elements {
			if !elm.IsCnf() {
				return false
			}
		}
		return true
	case "v":
		for _, elm := range c.elements {
			switch elm.(type) {
			case literal.Literal:
				continue
			case Clause:
				cl := elm.(Clause)
				if !cl.IsNegativeLiteral() {
					return false
				}
			}
		}
		return true
	}
	return false
}

// IsNegativeLiteral true if clause is of form ~A, but not ~(AvB)
func (c Clause) IsNegativeLiteral() bool {
	if c.Symbol == "~" {
		switch c.elements[0].(type) {
		case literal.Literal:
			return true
		default:
			return false
		}
	}
	return false
}
