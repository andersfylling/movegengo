package cmgg

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
