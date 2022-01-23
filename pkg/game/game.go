package game

import (
	"math/rand"
	//log "github.com/sirupsen/logrus"
)

type Game struct {
	board  Board
	size   int
	pieces []Piece
}

// A player that chooses a location from a list
type Player interface {
	Decide(Board, Piece, []Location) int
}

// A player that is responsible for choosing a location
// Nothing checks to make sure it is a valid location
// If a move is not possible, the player should return nil
type SmartPlayer interface {
	Move(Board, Piece) *Location
}

func New() *Game {
	size := DEFAULT_BOARD_SIZE
	b := make([][]bool, size)
	for i := range b {
		b[i] = make([]bool, size)
	}
	return NewFromBoard(b)
}

func NewFromBoard(b Board) *Game {
	//log.Info("Start to build game")
	size := len(b)
	g := Game{
		board:  b,
		size:   size,
		pieces: getPiecees(),
	}
	//log.Info("Finished building game")
	return &g
}

func (g *Game) pickRandomPiece() Piece {
	randomIndex := rand.Intn(len(g.pieces))
	return g.pieces[randomIndex]
}

func (g *Game) Play(p Player) (int, Piece) {

	smartPlayer, hasSmartPlayer := p.(SmartPlayer)
	if hasSmartPlayer {
		return g.PlaySmart(smartPlayer)
	}
	return g.play(func(nextPiece Piece) *Location {
		possibleLocations := g.board.GetPossibleLocations(nextPiece)
		if len(possibleLocations) == 0 {
			return nil
		}
		choiceIndex := p.Decide(g.board, nextPiece, possibleLocations)
		choice := possibleLocations[choiceIndex]
		return &choice
	})
}

func (g *Game) PlaySmart(p SmartPlayer) (int, Piece) {
	return g.play(func(nextPiece Piece) *Location {
		return p.Move(g.board, nextPiece)
	})
}

func (g *Game) play(chooseLocation func(Piece) *Location) (int, Piece) {
	score := 0
	for {
		//log.Info("Starting turn")
		nextPiece := g.pickRandomPiece()
		choice := chooseLocation(nextPiece)
		if choice == nil {
			return score, nextPiece
		}
		g.board.Place(nextPiece, *choice)
		score += 1
		//log.Info("Finished turn")
	}
}

func (g *Game) Show() {
	g.board.Show()
}
