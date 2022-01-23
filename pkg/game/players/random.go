package players

import (
	"math/rand"
	"time"

	"github.com/lukemassa/onezero/pkg/game"
)

type RandomPlayer struct {
	rand rand.Rand
}

func NewRandomPlayer() *RandomPlayer {
	r := RandomPlayer{
		rand: *rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return &r
}

func (r *RandomPlayer) Decide(b game.Board, p game.Piece, locations []game.Location) int {
	return r.rand.Intn(len(locations))
}
