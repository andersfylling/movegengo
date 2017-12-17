package cmgg

// MoveGen chess move generator
type MoveGen struct {
	moves  [255]uint16
	index  uint
	state  *GameState
	mover  *Move
	colour uint8
}

// NewMoveGen initiate a new MoveGen instance
func NewMoveGen() *MoveGen {
	return &MoveGen{index: 0, state: NewGameState(), mover: NewMove(0), colour: DefaultGameStateColour()}
}

// NewMoveGenByState creates a new movegen instance using the following state
func NewMoveGenByState(st *GameState) *MoveGen {
	return &MoveGen{index: 0, state: st, mover: NewMove(0), colour: (st.info & 0x10) >> 5}
}

// SetState update gamestate by reusing memory space
func (mg *MoveGen) SetState(st *GameState) {
	mg.state = st

	// reset any data associated with the old state
	mg.Clear()
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
func (mg *MoveGen) GetMove(index uint) uint16 {
	return mg.moves[index]
}

// CreateIterator Iterator pattern
func (mg *MoveGen) CreateIterator() *MoveGenIterator {
	return NewMoveGenIterator(0, mg.index, mg)
}

func (mg *MoveGen) isWhite() bool {
	return mg.colour == 1
}

// Moves are generated here
//

// GenerateMoves generates all the moves for an active player
func (mg *MoveGen) GenerateMoves() {
	if mg.isWhite() { // if the active player is white

	} else { // generate moves for the black player

	}

	mg.GeneratePawnMoves()
}

func (mg *MoveGen) GeneratePawnMoves() uint64 {
	pawns := mg.state.pieces[mg.colour*6]

	// attack left
	attacksLeft := mg.generatePawnLeftAttack(pawns)

	// attack right
	attacksRight := mg.generatePawnRightAttack(pawns)

	// single push
	singlePush := mg.generatePawnSinglePush()

	// double push
	mg.generatePawnDoublePush(singlePush)

	// results in possible attacks
	return attacksLeft | attacksRight
}

/**
 * Move all the pawns forward once and do a promotion check.
 * Promotions are handled, and removed from the resulting bitboard.
 *
 * @return A bitboard for all the pawns that moved forward, without being promoted.
 */
func (mg *MoveGen) generatePawnSinglePush() uint64 {
	pawns := mg.state.pieces[mg.colour*6]
	var to uint64

	if mg.isWhite() {
		to = (pawns << 8) &^ mg.state.colours[0] // ~(this->state.taken)
	} else {
		to = (pawns >> 8) &^ mg.state.colours[1] // ~(this->state.taken)
	}
	cache := to

	// remove promotion pieces
	to ^= mg.generatePromotions(0, to)

	// single push
	mg.mover.SetFlags(0)
	for i := LSB(to); i != 64; i = NLSB(&to, i) {
		if mg.isWhite() {
			mg.mover.SetFrom(uint16(i - 8))
		} else {
			mg.mover.SetFrom(uint16(i + 8))
		}
		mg.mover.SetTo(uint16(i))
		mg.moves[mg.index] = mg.mover.GetMove()
		mg.index++
	}

	return cache
} // end pawn generation

/**
 * Generate legal double push pawn moves.
 * This is a continuation on single pawn push, so the argument must be
 * the pawns that has already moved once for accurate results.
 *
 * @param pawns All pawn positions after single legal push (move). eg. 16711680ull
 * @return All the new pawn positions
 */
func (mg *MoveGen) generatePawnDoublePush(pawns uint64) uint64 {
	var to uint64

	if mg.isWhite() {
		to = ((pawns & 0xff0000) << 8) &^ mg.state.colours[0] // ~(this->state.taken)
	} else {
		to = ((pawns & 0xff0000000000) >> 8) &^ mg.state.colours[1] // ~(this->state.taken)
	}

	var attackDirection int
	if mg.isWhite() {
		attackDirection = -16
	} else {
		attackDirection = 16
	}

	cache := to
	mg.mover.SetFlags(1) // 0b0001, double push
	for i := LSB(to); i != 64; i = NLSB(&to, i) {
		mg.mover.SetFrom(uint16(i + attackDirection))
		mg.mover.SetTo(uint16(i))
		mg.moves[mg.index] = mg.mover.GetMove()
		mg.index++
	}

	return cache
}

/**
 * Generate all pawn attacks on the left side.
 * Promotions are handled as after movement.
 *
 * @param pawns uint64 over all pawns.
 * @return uint64 with positions reached.
 */
func (mg *MoveGen) generatePawnLeftAttack(pawns uint64) uint64 {
	area := uint64(0x7f7f7f7f7f7f7f7f)
	var attacks uint64
	if mg.isWhite() {
		attacks = (pawns & area) << 9
		attacks &= mg.state.colours[0] // TODO: en passant
	} else {
		attacks = (pawns & area) >> 7
		attacks &= mg.state.colours[1] // TODO: en passant
	}

	cache := attacks

	// promotions
	var attackDirection int
	if mg.isWhite() {
		attackDirection = 9
	} else {
		attackDirection = -7
	}
	attacks ^= mg.generatePromotions(attackDirection, attacks)

	mg.mover.SetFlags(4) // 0b0100, capture
	for i := LSB(attacks); i != 64; i = NLSB(&attacks, i) {
		mg.mover.SetFrom(uint16(i + attackDirection))
		mg.mover.SetTo(uint16(i))
		mg.moves[mg.index] = mg.mover.GetMove()
		mg.index++
	}

	return cache
}

/**
 * Generate all pawn attacks on the right side.
 * Promotions are handled as after movement.
 *
 * @param pawns uint64 over all pawns.
 * @return uint64 with positions reached.
 */
func (mg *MoveGen) generatePawnRightAttack(pawns uint64) uint64 {
	area := uint64(0xfefefefefefe00)
	var attacks uint64 // TODO: en passant
	if mg.isWhite() {
		attacks = (pawns & area) << 7
		attacks &= mg.state.colours[0]
	} else {
		attacks = (pawns & area) >> 9
		attacks &= mg.state.colours[1]
	}
	cache := attacks

	// promotions
	var attackDirection int
	if mg.isWhite() {
		attackDirection = -7
	} else {
		attackDirection = 9
	}
	attacks ^= mg.generatePromotions(attackDirection, attacks)

	// capture, 0b0100
	mg.mover.SetFlags(4)
	for i := LSB(attacks); i != 64; i = NLSB(&attacks, i) {
		mg.mover.SetFrom(uint16(i + attackDirection))
		mg.mover.SetTo(uint16(i))
		mg.moves[mg.index] = mg.mover.GetMove()
		mg.index++
	}

	return cache
}

/**
 * Generate promotion pieces.
 *
 * @tparam FROM the bitboard offset in int8, diff between from and to position index.
 * @param pawns
 * @return
 */
func (mg *MoveGen) generatePromotions(from int, pawns uint64) uint64 {
	promotions := pawns & 0xff000000000000ff

	// single push has a FROM of 1. since this is an offset.
	var flag uint8
	if from != 8 && from != -8 {
		flag = 12 // 0b1100
	} else {
		flag = 8 // 0b1000
	}

	for i := LSB(promotions); i != 64; i = NLSB(&promotions, i) {
		mg.mover.SetFrom(uint16(i + from))
		mg.mover.SetTo(uint16(i))

		var t uint8
		for ; t < 4; t++ {
			mg.mover.SetFlags(uint16(flag + t))
			mg.moves[mg.index] = mg.mover.GetMove()
			mg.index++
		}
	}

	return pawns & 0xff000000000000ff
}
