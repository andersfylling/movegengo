package cmgg

type GameState struct {
	colours [2]uint64
	pieces  [12]uint64
	info    uint8
	ep      uint8
}

func DefaultGameStateColour() uint8 {
	return 1
}

func NewGameState() *GameState {
	return &GameState{
		// default concated layouts
		colours: [2]uint64{
			BPieces,
			WPieces,
		},

		// Default layouts
		pieces: [12]uint64{
			BPawn,
			BRook,
			BKnight,
			BBishop,
			BQueen,
			BKing,

			WPawn,
			WRook,
			WKnight,
			WBishop,
			WQueen,
			WKing,
		},

		// castling: (5)0b11111, 4 == active player.
		// 0:blackkingcast, 1:whitekingcast, 2:blackqueencast, 3:whitequeencast, 4:white
		info: 15 | (DefaultGameStateColour() << 5),

		// en passant position. 0 means no en passant
		ep: 0,
	}
}

// Print represent the content as a chess board
func (state *GameState) Print() {
	// TODO
}
