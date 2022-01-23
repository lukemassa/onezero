package players

import (
	"github.com/lukemassa/onezero/pkg/game"
)

type UpperLeftPlayer struct{}

func (r *UpperLeftPlayer) Decide(b game.Board, p game.Piece, locations []game.Location) int {
	return 0
}
