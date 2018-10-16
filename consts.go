package movegengo

// MaxMoves per round
const MaxMoves = 256

// move constants
//

// FlagPromotion promotion bit
const FlagPromotion uint16 = 32768 // 0b1000000000000000
// FlagCapture capture bit
const FlagCapture uint16 = 16384 // 0b0100000000000000
// FlagSpecial1 special1 bit
const FlagSpecial1 uint16 = 8192 // 0b0010000000000000
// FlagSpecial0 special2 bit
const FlagSpecial0 uint16 = 4096 // 0b0001000000000000
// RangeFlag bit range of all bit positions used for flags
const RangeFlag uint16 = 61440 // 0b1111000000000000
// RangeFrom bit range which holds the From position as a uint8 position
const RangeFrom uint16 = 4032 // 0b0000111111000000
// RangeTo bit range which holds the To position as a uint8 position
const RangeTo uint16 = 63 // 0b0000000000111111

// piece index

// Bitboards default positions
//

// For all the different black pieces
//

// BBishop black bishops
const BBishop uint64 = 2594073385365405696

// BKing black king
const BKing uint64 = 576460752303423488

// BKnight black knight
const BKnight uint64 = 4755801206503243776

// BPawn black pawn
const BPawn uint64 = 71776119061217280

// BQueen black queen
const BQueen uint64 = 1152921504606846976

// BRook black rooks
const BRook uint64 = 9295429630892703744

// BPieces all the black pieces concated
const BPieces uint64 = BBishop | BKnight | BPawn | BQueen | BKing | BRook

// For all the different white pieces
//

//WBishop white bishops
const WBishop uint64 = 36

// WKnight white knights
const WKnight uint64 = 66

// WPawn white pawns
const WPawn uint64 = 65280

// WQueen white queen
const WQueen uint64 = 16

// WKing white king
const WKing uint64 = 8

// WRook white rooks
const WRook uint64 = 129

// WPieces all the white pieces concated
const WPieces uint64 = WBishop | WKnight | WPawn | WQueen | WKing | WRook
