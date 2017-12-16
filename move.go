package cmgg

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

func (m *Move) SetMoveFromInstance(m2 *Move) {
	m.move = m2.GetMove()
}

func (m *Move) SetMove(move uint16) {
	m.move = move
}

func (m *Move) GetMove() uint16 {
	return m.move
}

func (m *Move) GetTo() uint16 {
	return m.move & RangeTo
}

func (m *Move) GetFrom() uint16 {
	return (m.move & RangeFrom) >> 6
}
func (m *Move) GetFlags() uint16 {
	return (m.move & RangeFlag) >> 12
}

func (m *Move) SetTo(to uint16) {
	m.move &= ^RangeTo
	m.move |= to & RangeTo
}
func (m *Move) SetFrom(from uint16) {
	m.move &= ^RangeFrom
	m.move |= (from << 6) & RangeFrom // why isn't from here just rendered [0, 3]?
}
func (m *Move) SetFlags(flags uint16) {
	m.move &= ^RangeFlag
	m.move |= (flags << 12) & RangeFlag
}

func (m *Move) HasPromotion() bool {
	return (m.move & FlagPromotion) != 0
}

func (m *Move) HasCapture() bool {
	return (m.move & FlagCapture) != 0
}

func (m *Move) HasSpecial1() bool {
	return (m.move & FlagSpecial1) != 0
}

func (m *Move) hasSpecial0() bool {
	return (m.move & FlagSpecial0) != 0
}

func (m *Move) IsQuietMoves() bool {
	return (m.move & RangeFlag) == 0
}

func (m *Move) IsDoublePawnPush() bool {
	return ((m.move & RangeFlag) == FlagSpecial0)
}

func (m *Move) IsKingCastle() bool {
	return ((m.move & RangeFlag) == FlagSpecial1)
}

func (m *Move) IsQueenCastle() bool {
	return ((m.move & RangeFlag) == (FlagSpecial0 | FlagSpecial1))
}

func (m *Move) IsEPCapture() bool {
	return ((m.move & RangeFlag) == (FlagSpecial0 | FlagCapture))
}

/**
 * When the moving piece has caused check.
 * @return true on check
 */
func (m *Move) IsCheck() bool {
	return (m.move & RangeFlag) == (FlagCapture | FlagSpecial1 | FlagSpecial0)
}

func (m *Move) IsKnightPromotion() bool {
	return ((m.move & RangeFlag) == FlagPromotion)
}

func (m *Move) IsBishopPromotion() bool {
	return ((m.move & RangeFlag) == (FlagPromotion | FlagSpecial0))
}

func (m *Move) IsRookPromotion() bool {
	return ((m.move & RangeFlag) == (FlagPromotion | FlagSpecial1))
}

func (m *Move) IsQueenPromotion() bool {
	return ((m.move & RangeFlag) == (FlagPromotion | FlagSpecial1 | FlagSpecial0))
}

func (m *Move) IsKnightPromoCapture() bool {
	return ((m.move & RangeFlag) == (FlagPromotion | FlagCapture))
}

func (m *Move) IsBishopPromoCapture() bool {
	return ((m.move & RangeFlag) == (FlagPromotion | FlagCapture | FlagSpecial0))
}

func (m *Move) IsRookPromoCapture() bool {
	return ((m.move & RangeFlag) == (FlagPromotion | FlagCapture | FlagSpecial1))
}

func (m *Move) IsQueenPromoCapture() bool {
	return ((m.move & RangeFlag) == (FlagPromotion | FlagCapture | FlagSpecial1 | FlagSpecial0))
}

func (m *Move) GetButterflyIndex() uint16 {
	return m.move & (RangeFrom | RangeTo)
}
func (m *Move) Equal(a *Move) bool {
	return m.move == a.GetMove()
}
func (m *Move) Not(a *Move) bool {
	return m.move != a.GetMove()
}
