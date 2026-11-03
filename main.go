package main

import (
	"fmt"

	"github.com/corentings/chess/v2"
)

var game = chess.NewGame()

func main() {

	fmt.Println("test")
	move_piece("e4")
	fmt.Println(game.Position().Board().Draw())
}

func move_piece(piece string) {
	err := game.PushNotationMove(piece, chess.AlgebraicNotation{}, nil)

	if err != nil {
		panic(err)
	}
}
