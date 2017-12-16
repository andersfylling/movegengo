package cmgg

// MoveGenIterator chess move generator
type MoveGenIterator struct {
	index int
	size  int
	mg    *MoveGen
}

// NewMoveGenIterator Creates an iterator for MoveGen
func NewMoveGenIterator(i, size int, mg *MoveGen) *MoveGenIterator {
	return &MoveGenIterator{index: i, size: size, mg: mg}
}

// Iterator pattern
// Begin(), End(), Next(), GetMove(uint8)

// Begin [iterator] return the first element
func (it *MoveGenIterator) Begin() int {
	return 0
}

// End [iterator] return the element after the last
func (it *MoveGenIterator) End() int {
	return it.size
}

// Good checks if the iteration is within bounds
func (it *MoveGenIterator) Good() bool {
	return it.index < it.size
}

// Next [iterator] get the next index
func (it *MoveGenIterator) Next() int {
	it.index++

	return it.index
}

// GetIndex returns the current iterator index (progress)
func (it *MoveGenIterator) GetIndex() int {
	return it.index
}

// GetMove explicit to MoveGen and returns an encoded chess move
func (it *MoveGenIterator) GetMove() uint16 {
	return it.mg.GetMove(it.index)
}

// Verify it legality of Iterator interface implementation
var _ Iterator = (*MoveGenIterator)(nil)
