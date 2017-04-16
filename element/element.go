package element

// Element base type
type Element interface {
	String() string
	IsCnf() bool
}
