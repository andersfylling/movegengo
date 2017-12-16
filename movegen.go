package cmgg

// MoveGen chess move generator
type MoveGen struct {
	moves [255]uint16
	index uint
}

// NewMoveGen initiate a new MoveGen instance
func NewMoveGen() *MoveGen {
	return &MoveGen{index: 0}
}

// Clear the moves list tracker & iterator
func (mg *MoveGen) Clear() {
	mg.index = 0
}

// Size get the number of stored moves so far
func (mg *MoveGen) Size() uint {
	return mg.index
}

// AddMove to the stack
func (mg *MoveGen) AddMove(move uint16) {
	mg.moves[mg.index] = move
	mg.index++
}

// SetMove set a move at a specific location
func (mg *MoveGen) SetMove(move uint16, index int) {
	mg.moves[index] = move
}

// GetMove returns the move for a given index
func (mg *MoveGen) GetMove(index int) uint16 {
	return mg.moves[index]
}

// CreateIterator Iterator pattern
func (mg *MoveGen) CreateIterator() *MoveGenIterator {
	return NewMoveGenIterator(0, len(mg.moves), mg)
}
