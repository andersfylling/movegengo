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

## Architecture ##
Currently the goal is to reach a working move generator. I'll then record a bunch of input-output data, and use that as a control for optimizations later on. Right now the idea is to create 16bit encoded moves and only produce legal moves from the movegenerator. Not generating every psuedo move and validate them using a function.

The features ment to be accessed by other developers, mainly, are placed in the root folder. And both package and file by feature design is used.

## Bitboards ##
I know there are different bitboard version, as in the order of the chess board squares. But this model aims at 0 == F1 and 63 == A8, from the white perspective. I do want the possibility to swap between bitboard layouts later, but as I'm worried about the performance hit of not having a precompiler, implementations for different bitboard layouts will most likely gain their own repositories (if this ever comes that far).

## Contributing ##
In your PR (Pull Request) explain what you have added/implemented or changed. If it affects performance, or you are asked to provide perfomance information, do a control test from the original branch and your new branch using perft & benchmarks. Any new feature implemented <b>must</b> have unit tests, and possibly benchmarks if required.

## Requests and bugs ##
I'm flexible when it comes to new ideas. I even want to try using a 32bit encoded move instead as the newer CPU's from AMD has way more L1 cache. But that's for a later consideration. Use the issue tracker for everything; questions, bugs, requests, concerns, etc.
