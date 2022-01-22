package players

import (
	"math/rand"

	"github.com/lukemassa/onezero/pkg/game"
)

type RandomPlayer struct{}

func (r RandomPlayer) Decide(b game.Board, p game.Piece, locations []game.Location) int {
	return rand.Intn(len(locations))
}
