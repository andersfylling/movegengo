package movegengo

import "strconv"

// Move is used to hold a move and manipulate or extract data from it
type Move struct {
	move uint16
}

// NewMove Creates a new move from a encoded move
func NewMove(move uint16) *Move {
	return &Move{move: move}
}

// NewMoveDetail creates a new encoded move based on primitive values
func NewMoveDetail(from, to, flags uint16) *Move {
	return &Move{move: (uint16(flags&0xf) << 12) | (uint16(from&0x3f) << 6) | uint16(to&0x3f)}
}

// SetMoveFromInstance ...
func (m *Move) SetMoveFromInstance(m2 *Move) {
	m.move = m2.GetMove()
}

// SetMove ...
func (m *Move) SetMove(move uint16) {
	m.move = move
}

// GetMove ...
func (m *Move) GetMove() uint16 {
	return m.move
}

// To ...
func (m *Move) To() uint16 {
	return m.move & RangeTo
}

// From ...
func (m *Move) From() uint16 {
	return (m.move & RangeFrom) >> 6
}

// Flags ...
func (m *Move) Flags() uint16 {
	return (m.move & RangeFlag) >> 12
}

// SetTo ...
func (m *Move) SetTo(to uint16) {
	m.move &= ^RangeTo
	m.move |= to & RangeTo
}

// SetFrom ...
func (m *Move) SetFrom(from uint16) {
	m.move &= ^RangeFrom
	m.move |= (from << 6) & RangeFrom // why isn't from here just rendered [0, 3]?
}

// SetFlags ...
func (m *Move) SetFlags(flags uint16) {
	m.move &= ^RangeFlag
	m.move |= (flags << 12) & RangeFlag
}

// HasPromotion ...
func (m *Move) HasPromotion() bool {
	return (m.move & FlagPromotion) != 0
}

// HasCapture ...
func (m *Move) HasCapture() bool {
	return (m.move & FlagCapture) != 0
}

// HasSpecial1 ...
func (m *Move) HasSpecial1() bool {
	return (m.move & FlagSpecial1) != 0
}

// hasSpecial0 ...
func (m *Move) hasSpecial0() bool {
	return (m.move & FlagSpecial0) != 0
}

// IsQuietMoves ...
func (m *Move) IsQuietMoves() bool {
	return (m.move & RangeFlag) == 0
}

// IsDoublePawnPush ...
func (m *Move) IsDoublePawnPush() bool {
	return (m.move & RangeFlag) == FlagSpecial0
}

// IsKingCastle ...
func (m *Move) IsKingCastle() bool {
	return (m.move & RangeFlag) == FlagSpecial1
}

// IsQueenCastle ...
func (m *Move) IsQueenCastle() bool {
	return (m.move & RangeFlag) == (FlagSpecial0 | FlagSpecial1)
}

// IsEPCapture ...
func (m *Move) IsEPCapture() bool {
	return (m.move & RangeFlag) == (FlagSpecial0 | FlagCapture)
}

// IsCheck When the moving piece has caused check. Returns true on check.
func (m *Move) IsCheck() bool {
	return (m.move & RangeFlag) == (FlagCapture | FlagSpecial1 | FlagSpecial0)
}

// IsKnightPromotion ...
func (m *Move) IsKnightPromotion() bool {
	return (m.move & RangeFlag) == FlagPromotion
}

// IsBishopPromotion ...
func (m *Move) IsBishopPromotion() bool {
	return (m.move & RangeFlag) == (FlagPromotion | FlagSpecial0)
}

// IsRookPromotion ...
func (m *Move) IsRookPromotion() bool {
	return (m.move & RangeFlag) == (FlagPromotion | FlagSpecial1)
}

// IsQueenPromotion ...
func (m *Move) IsQueenPromotion() bool {
	return (m.move & RangeFlag) == (FlagPromotion | FlagSpecial1 | FlagSpecial0)
}

// IsKnightPromoCapture ...
func (m *Move) IsKnightPromoCapture() bool {
	return (m.move & RangeFlag) == (FlagPromotion | FlagCapture)
}

// IsBishopPromoCapture ...
func (m *Move) IsBishopPromoCapture() bool {
	return (m.move & RangeFlag) == (FlagPromotion | FlagCapture | FlagSpecial0)
}

// IsRookPromoCapture ...
func (m *Move) IsRookPromoCapture() bool {
	return (m.move & RangeFlag) == (FlagPromotion | FlagCapture | FlagSpecial1)
}

// IsQueenPromoCapture ...
func (m *Move) IsQueenPromoCapture() bool {
	return (m.move & RangeFlag) == (FlagPromotion | FlagCapture | FlagSpecial1 | FlagSpecial0)
}

// ButterflyIndex ...
func (m *Move) ButterflyIndex() uint16 {
	return m.move & (RangeFrom | RangeTo)
}

// Equal ...
func (m *Move) Equal(a *Move) bool {
	return m.move == a.GetMove()
}

// Not ...
func (m *Move) Not(a *Move) bool {
	return m.move != a.GetMove()
}

// String shows from and to. Essentially what piece was moved, from where it was moved and it's new position
func (m *Move) String() string {
	return "From: " + string(strconv.Itoa(int(m.From()))) + ", To: " + string(strconv.Itoa(int(m.To())))
}
