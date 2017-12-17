package cmgg

import (
	"strconv"
	"testing"
)

func itos(i int) string {
	return string(strconv.Itoa(i))
}
func u16tos(i uint16) string {
	return string(strconv.Itoa(int(i)))
}

func TestMoveGenIterator(t *testing.T) {
	mg := NewMoveGen()
	mg.GenerateMoves()

	it := mg.CreateIterator()
	if it.Begin() != 0 {
		t.Errorf("Begin() returned %d, wants 0", it.Begin())
	}
	// TODO

	// should be possible to check if a uint8 is smaller than a int... really.
	//if it.End() > mg.Size() {
	//	t.Errorf("End() was larger than MoveGen.index. Got %d, wants >%d", it.End(), mg.Size())
	//}
	it.End() //...

	// iterate
	for ; it.Good(); it.Next() {
		//println("move{" + itos(it.GetIndex()) + "}: " + u16tos(it.GetMove()))
		it.GetIndex()
		it.GetMove()
	}
}
