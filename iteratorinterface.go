package cmgg

// Iterator has the bare minimum to implement an iterator pattern
type Iterator interface {
	Begin() int
	End() int
	Next() int
	Good() bool
}
