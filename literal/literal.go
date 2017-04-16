package literal

// Literal aka "A", "B", ...
type Literal struct {
	Name string
}

func (l Literal) String() string {
	return l.Name
}

// IsCnf always true for literal
func (l Literal) IsCnf() bool {
	return true
}
