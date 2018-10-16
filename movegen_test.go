package movegengo

import "testing"

func TestMoveGenSize(t *testing.T) {
	movegen := NewMoveGen()

	if movegen.Size() != 0 {
		t.Error("Error: MoveGen size should be 0 at initialized instance.")
	}

	// increase Size
	movegen.AddMove(453)
	if movegen.Size() != 1 {
		t.Error("Error: MoveGen size should be 1 after adding one move.")
	}

	// clear move list
	movegen.Clear()
	if movegen.Size() != 0 {
		t.Error("Error: MoveGen size should be 0 after clearing move list.")
	}

	// setting a move, should not affect the size. at all.
	movegen.SetMove(432, 6)
	if movegen.Size() != 0 {
		t.Error("Error: MoveGen size should not be affected by updating a move.")
	}
}

func TestSinglePawnPush(t *testing.T) {
	movegen := NewMoveGen()
	movegen.generatePawnSinglePush()

	wants := []uint16{
		528, 593, 658, 723, 788, 853, 918, 983,
	}

	i := 0
	for it := movegen.CreateIterator(); it.Good() && i < len(wants); it.Next() {
		if it.GetMove() != wants[i] {
			t.Errorf("SinglePawnPush error: got %d, wants %d", it.GetMove(), wants[i])
		}
		i++
	}
}

func TestDoublePawnPush(t *testing.T) {
	movegen := NewMoveGen()
	movegen.generatePawnDoublePush(16711680)

	wants := []uint16{
		4632, 4697, 4762, 4827, 4892, 4957, 5022, 5087,
	}

	i := 0
	for it := movegen.CreateIterator(); it.Good() && i < len(wants); it.Next() {
		if it.GetMove() != wants[i] {
			t.Errorf("SinglePawnPush error: got %d, wants %d", it.GetMove(), wants[i])
		}
		i++
	}

	movegen.Clear()
	movegen.generatePawnDoublePush(16711680 << 8)
	if movegen.Size() != 0 {
		t.Error("MoveGen generated double pawn push without pawns in legal region")
	}
}

func TestPawnAttackRight(t *testing.T) {
	movegen := NewMoveGen()

	// try with no hostiles in bound
	movegen.generatePawnRightAttack(16711680)
	if movegen.Size() != 0 {
		t.Error("MoveGen generated pawn attacks(right) without any hostiles in sight")
	}

	// put all the pawns on rank 6, with hostiles on full rank 7
	movegen.generatePawnRightAttack(0xff0000000000)
	if movegen.Size() != 7 {
		t.Error("MoveGen generated more pawn attacks(right) than hostiles in sight")
	}

	//TODO en passant
}

func TestPawnAttackLeft(t *testing.T) {
	movegen := NewMoveGen()

	// try with no hostiles in bound
	movegen.generatePawnLeftAttack(16711680)
	if movegen.Size() != 0 {
		t.Error("MoveGen generated pawn attacks(left) without any hostiles in sight")
	}

	// put all the pawns on rank 6, with hostiles on full rank 7
	movegen.generatePawnLeftAttack(0xff0000000000)
	if movegen.Size() != 7 {
		t.Error("MoveGen generated more pawn attacks(left) than hostiles in sight")
	}

	//TODO en passant
}

func TestPawnPromotion(t *testing.T) {
	movegen := NewMoveGen()

	// try with no promotions
	movegen.generatePromotions(0, 16711680)
	if movegen.Size() != 0 {
		t.Error("MoveGen generated pawn pormotions without being close to the end")
	}

	// should not be a promotion by capture
	movegen.generatePromotions(-8, 0x4000000000000000)
	for it := movegen.CreateIterator(); it.Good(); it.Next() {
		move := NewMove(it.GetMove())
		if move.Flags() < 8 {
			t.Errorf("PawnPromotion error: flag was below 8. Got %d", move.Flags())
		}
		if move.Flags() >= 12 {
			t.Errorf("PawnPromotion error: flag was contained capture flag. Got %d", move.Flags())
		}
	}

	// promotion by capture
	movegen.Clear()
	movegen.generatePromotions(-7, 0x4000000000000000)
	for it := movegen.CreateIterator(); it.Good(); it.Next() {
		move := NewMove(it.GetMove())
		if move.Flags() < 12 {
			t.Errorf("PawnPromotion error: Did not contain capture flag (right). Got %d", move.Flags())
		}
	}

	// promotion by capture
	movegen.Clear()
	movegen.generatePromotions(-9, 0x4000000000000000)
	for it := movegen.CreateIterator(); it.Good(); it.Next() {
		move := NewMove(it.GetMove())
		if move.Flags() < 8 {
			t.Errorf("PawnPromotion error: Did not contain capture flag (left). Got %d", move.Flags())
		}
	}
}

func TestKnightMove(t *testing.T) {
	//movegen := NewMoveGen()

	//if movegen.generateKnightMove(0x1000000000) != 0x28440044280000 {
	//	t.Error("KnightMove error: attacks for 0x1000000000 is wrong.")
	//}
}

func generateKnightMoves(x uint64) uint64 {
	attacks := uint64(0x28440044280000)
	mask := uint64(0x7c7c7c7c7c0000)

	return mask + attacks
}

// create all knight moves
func generateKnightTable() []uint64 {
	return []uint64{0}
}

func BenchmarkKnight(b *testing.B) {

	for n := 0; n < b.N; n++ {

	}
}
