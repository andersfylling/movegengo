package cmgg

// Iterator has the bare minimum to implement an iterator pattern
type Iterator interface {
	Begin() uint
	End() uint
	Next() uint
	Good() bool
}
