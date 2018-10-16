# MovegenGo [![Documentation](https://godoc.org/github.com/chessmodule/movegengo?status.svg)](http://godoc.org/github.com/chessmodule/movegengo)

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/for-you.svg)](https://forthebadge.com)

## Health
| Branch       | Build status  | Code climate | Go Report Card | Code Coverage |
| ------------ |:-------------:|:---------------:|:-------------:|:----------------:|
| master     | [![CircleCI](https://circleci.com/gh/chessmodule/movegengo/tree/master.svg?style=shield)](https://circleci.com/gh/chessmodule/movegengo/tree/master) | [![Maintainability](https://api.codeclimate.com/v1/badges/f28b832369e4027522e7/maintainability)](https://codeclimate.com/github/chessmodule/movegengo/maintainability) | [![Go Report Card](https://goreportcard.com/badge/github.com/chessmodule/movegengo)](https://goreportcard.com/report/github.com/chessmodule/movegengo) | [![Test Coverage](https://api.codeclimate.com/v1/badges/f28b832369e4027522e7/test_coverage)](https://codeclimate.com/github/chessmodule/movegengo/test_coverage) |

## About
This package is a chess move generator that encodes moves. It uses a hybrid of psuedo and legal moves. It has implemented state pattern (using the game state) and iterator pattern (iterating over the move list). I will be implementing a array pointer with an range as argument to directly store the moves into a stack/game tree for a speed improvement.

This package also includes a Move "class" for encoding and decoding moves and move tables for certain pieces like knights.

## Bitboard layout
```

# Chess board layout

63 62 61 60 59 58 57 56  | 8
55 54 53 52 51 50 49 48  | 7
47 46 45 44 43 42 41 40  | 6
39 38 37 36 35 34 33 32  | 5
31 30 29 28 27 26 25 24  | 4
23 22 21 20 19 18 17 16  | 3
15 14 13 12 11 10 09 08  | 2
07 06 05 04 03 02 01 00  | 1
_________________________|
 A  B  C  D  E  F  G  H

# uint64 layout by index
A8 B8 C8 D8 ... E1 F1 G1 H1
63 62 61 60 ... 03 02 01 00

```
The position H1 is found at `bitboard & 0x1`, while A8 is found at `bitboard & (0x1<<63)`.

## Quick start using the iterator
```go
package main

import mg "github.com/chessmodule/movegengo"
import "fmt"

func main() {
  // For a move generator with a default game state use:
  //   movegen := mg.NewMoveGen()
  // If you have a populated game state (mg.GameState) use:
  //   movegen := mg.NewMoveGenByState(gs)
  movegen := mg.NewMoveGen()

  movegen.GenerateMoves() // generates all the moves
  for it := movegen.CreateIterator(); it.Good(); it.Next() {
    mover := mg.NewMove(it.GetMove()) //GetMove returns a uint16 encoded move
    fmt.Println("move: " + mover.ToStr()) // shows from and to values
  }
}
```