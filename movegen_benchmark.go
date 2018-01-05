package cmgg

import "testing"

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
