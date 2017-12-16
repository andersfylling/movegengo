package cmgg

import "testing"

func requireEqual(t *testing.T, a, b uint16) bool {
	if a != b {
		t.Errorf("Error: missmatch. Got %d, wants %d", a, b)
		return false
	}

	return true
}

func requireTrue(t *testing.T, a bool) bool {
	if !a {
		t.Error("Error: Was not true. Got false, wants true")
		return false
	}

	return true
}

func requireNotEqual(t *testing.T, a, b uint16) bool {
	if a == b {
		t.Errorf("Error: matched, these aren't supposed to equal. Got %d, wants %d", a, b)
		return false
	}

	return true
}

func TestToAndFrom(t *testing.T) {
	m := NewMove(0)

	m.SetFrom(8)
	requireEqual(t, m.GetFrom(), 8)
	requireEqual(t, m.GetTo(), 0)
	requireEqual(t, m.GetFlags(), 0)

	m.SetFrom(8)
	requireEqual(t, m.GetFrom(), 8)
	requireEqual(t, m.GetTo(), 0)
	requireEqual(t, m.GetFlags(), 0)

	m.SetTo(16)
	requireEqual(t, m.GetFrom(), 8)
	requireEqual(t, m.GetTo(), 16)
	requireEqual(t, m.GetFlags(), 0)

	m.SetTo(0)
	requireEqual(t, m.GetFrom(), 8)
	requireEqual(t, m.GetTo(), 0)
	requireEqual(t, m.GetFlags(), 0)

	m.SetTo(24)
	requireEqual(t, m.GetFrom(), 8)
	requireEqual(t, m.GetTo(), 24)
	requireEqual(t, m.GetFlags(), 0)

	m.SetFrom(0)
	requireEqual(t, m.GetFrom(), 0)
	requireEqual(t, m.GetTo(), 24)
	requireEqual(t, m.GetFlags(), 0)

	m.SetTo(3)
	requireEqual(t, m.GetFrom(), 0)
	requireEqual(t, m.GetTo(), 3)
	requireEqual(t, m.GetFlags(), 0)

	// move using constructor
	m2 := NewMoveDetail(8, 24, 0)
	requireEqual(t, m2.GetFrom(), 8)
	requireEqual(t, m2.GetTo(), 24)
	requireEqual(t, m2.GetFlags(), 0)

}

func TestFlag(t *testing.T) {
	//flags
	m3 := NewMoveDetail(8, 24, 15)
	requireEqual(t, m3.GetFrom(), 8)
	requireEqual(t, m3.GetTo(), 24)
	requireEqual(t, m3.GetFlags(), 15)

	m3.SetFlags(7)
	requireEqual(t, m3.GetFrom(), 8)
	requireEqual(t, m3.GetTo(), 24)
	requireEqual(t, m3.GetFlags(), 7)

	m3.SetFlags(0)
	requireEqual(t, m3.GetFrom(), 8)
	requireEqual(t, m3.GetTo(), 24)
	requireEqual(t, m3.GetFlags(), 0)

	m3.SetFlags(7)
	m3.SetFrom(4)
	m3.SetTo(62)
	requireEqual(t, m3.GetFrom(), 4)
	requireEqual(t, m3.GetTo(), 62)
	requireEqual(t, m3.GetFlags(), 7)
}

func TestComparison(t *testing.T) {
	move := NewMove(0)
	move.SetFrom(8)
	move.SetTo(24)

	// move using constructor
	m2 := NewMoveDetail(8, 24, 0)

	//flags
	m3 := NewMoveDetail(8, 24, 15)

	m3.SetFlags(7)
	m3.SetFrom(4)
	m3.SetTo(62)
	requireEqual(t, m3.GetFrom(), 4)
	requireEqual(t, m3.GetTo(), 62)
	requireEqual(t, m3.GetFlags(), 7)

	requireNotEqual(t, m2.GetMove(), m3.GetMove())
	requireNotEqual(t, m2.GetButterflyIndex(), m3.GetButterflyIndex())

	m2.SetMoveFromInstance(m3)
	requireEqual(t, m2.GetMove(), m3.GetMove())
	requireEqual(t, m2.GetFrom(), 4)
	requireEqual(t, m2.GetTo(), 62)
	requireEqual(t, m2.GetFlags(), 7)

	m3.SetMove(move.GetMove())
	requireEqual(t, m3.GetMove(), move.GetMove())
	requireTrue(t, m3.Equal(move))
	requireNotEqual(t, m2.GetMove(), move.GetMove())
	requireTrue(t, m2.Not(move))

	m2.SetMove(move.GetMove())
	requireEqual(t, m2.GetMove(), move.GetMove())
}

func TestAttacks(t *testing.T) {
	move := NewMove(0)

	// check all flags
	var i uint16
	for ; i <= 15; i++ {
		move.SetFlags(i)
		requireEqual(t, move.GetFlags(), i)
	}

	i = 0

	// quiet move
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsQuietMoves())

	// double pawn push
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsDoublePawnPush())

	// king castle
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsKingCastle())

	// queen castle
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsQueenCastle())

	// captures
	move.SetFlags(i)
	i++
	requireTrue(t, move.HasCapture())

	// ep-capture
	move.SetFlags(i)
	i++ // 5
	requireTrue(t, move.IsEPCapture())
	// 6 and 7 isn't used by chessprogramming.com
	i += 2

	// knight-promotion
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsKnightPromotion())

	// bishop-promotion
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsBishopPromotion())

	// rook-promotion
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsRookPromotion())

	// queen-promotion
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsQueenPromotion())

	// knight-promo capture
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsKnightPromoCapture())

	// bishop-promo capture
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsBishopPromoCapture())

	// rook-promo capture
	move.SetFlags(i)
	i++
	requireTrue(t, move.IsRookPromoCapture())

	// queen-promo capture
	move.SetFlags(i)
	requireTrue(t, move.IsQueenPromoCapture())

	// custom flags

	// check
	move.SetFlags(7)
	requireTrue(t, move.IsCheck()) // when the moving piece has caused a check
}
