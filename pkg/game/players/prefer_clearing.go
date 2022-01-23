package players

import (
	"github.com/lukemassa/onezero/pkg/game"
)

type PreferClearingPlayer struct{}

func (r PreferClearingPlayer) Decide(b game.Board, p game.Piece, locations []game.Location) int {
	// Return the first option that clears, default to 0

	size := len(b)
	for i := 0; i < len(locations); i++ {
		newBoard := game.NewBlankBoard()
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				newBoard[x][y] = b[x][y]
			}
		}
		location := locations[i]
		// If placing here clears, then choose it
		//fmt.Println("***********************************")
		//b.Show()
		cleared := newBoard.Place(p, location)
		//		fmt.Println(cleared)
		if cleared {
			// fmt.Printf("Found clearing: %v\n", location)
			//	b.Show()
			//	p.Show()
			return i
		}

	}
	return 0
}
