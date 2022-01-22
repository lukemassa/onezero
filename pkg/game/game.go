package game

import (
	"math/rand"

	log "github.com/sirupsen/logrus"
)

type Game struct {
	board  Board
	size   int
	pieces []Piece
	player Player
}

type Player interface {
	Decide(Board, Piece, []Location) int
}

func New(p Player) *Game {
	log.Info("Start to build game")
	size := DEFAULT_BOARD_SIZE
	b := make([][]bool, size)
	for i := range b {
		b[i] = make([]bool, size)
	}
	g := Game{
		board:  b,
		size:   size,
		pieces: getPiecees(),
		player: p,
	}
	log.Info("Finished building game")
	return &g
}

func (g *Game) pickRandomPiece() Piece {
	randomIndex := rand.Intn(len(g.pieces))
	return g.pieces[randomIndex]
}

func (g *Game) Play() (int, Piece) {
	score := 0
	for {
		log.Info("Starting turn")
		nextPiece := g.pickRandomPiece()
		possibleLocations := g.board.getPossibleLocations(nextPiece)
		if len(possibleLocations) == 0 {
			return score, nextPiece
		}
		choiceIndex := g.player.Decide(g.board, nextPiece, possibleLocations)
		choice := possibleLocations[choiceIndex]
		g.board.place(nextPiece, choice)
		score += 1
		log.Info("Finished turn")
	}
}

func (g *Game) Show() {
	g.board.Show()
}