package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/corentings/chess/v2"
	"github.com/corentings/chess/v2/uci"
	chessreview "github.com/ksysoev/chess-review"
)

var game = chess.NewGame()
var eng *uci.Engine

func main() {
	reviewer, err := chessreview.New("/usr/games/stockfish",
		chessreview.WithDepth(10))
	if err != nil {
		log.Fatal(err)
	}

	eng, err = uci.New("/usr/games/stockfish")
	if err != nil {
		log.Fatal(err)
	}

	err = eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdSetOption{
		Name:  "Skill Level",
		Value: "20",
	}, uci.CmdUCINewGame)

	if err != nil {
		log.Fatal(err)
	}

	stockFishMove()
	movePiece("e5")
	stockFishMove()

	fmt.Println(game.Position().Board().Draw())
	review, err := reviewer.ReviewGame(context.Background(), game.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(review[0].ScoreDelta)
	fmt.Println(review[1].ScoreDelta)
	fmt.Println(review[2].ScoreDelta)
	reviewer.Close()
}

func movePiece(piece string) {
	err := game.PushNotationMove(piece, chess.AlgebraicNotation{}, nil)

	if err != nil {
		panic(err)
	}
}
func stockFishMove() {
	cmdGo := uci.CmdGo{MoveTime: time.Second * 5}
	cmdPos := uci.CmdPosition{Position: game.Position()}
	err := eng.Run(cmdPos, cmdGo)
	if err != nil {
		log.Fatal(err)
	}
	stockfishMove := eng.SearchResults().BestMove
	movePiece(stockfishMove.String())

}
