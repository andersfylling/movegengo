# [cmgg] Chess Move Generator in GoLang #
This package is a chess move generator that encodes moves. It uses a hybrid of psuedo and legal moves. It has implemented state pattern (using the game state) and iterator pattern (iterating over the move list). I will be implementing a array pointer with an range as argument to directly store the moves into a stack/game tree for a speed improvement.

This package also includes a Move "class" for encoding and decoding moves and move tables for certain pieces like knights.

## Usage ##
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
