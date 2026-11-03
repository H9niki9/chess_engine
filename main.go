package main

import (
	"context"
	"fmt"
	"log"

	"github.com/corentings/chess/v2"
	chessreview "github.com/ksysoev/chess-review"
)

var game = chess.NewGame()

func main() {
	reviewer, err := chessreview.New("/usr/games/stockfish")
	if err != nil {
		log.Fatal(err)
	}
	move_piece("e4")
	fmt.Println(game.Position().Board().Draw())
	review, err := reviewer.ReviewGame(context.Background(), game.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(review)
	reviewer.Close()
}

func move_piece(piece string) {
	err := game.PushNotationMove(piece, chess.AlgebraicNotation{}, nil)

	if err != nil {
		panic(err)
	}
}
