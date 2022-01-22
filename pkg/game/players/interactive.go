package players

import (
	"fmt"
	"strconv"

	"github.com/lukemassa/onezero/pkg/game"
)

type InteractivePlayer struct{}

func (i InteractivePlayer) Decide(b game.Board, p game.Piece, locations []game.Location) int {
	fmt.Print("\033[H\033[2J")
	b.Show()
	p.Show()
	fmt.Println("Options:")
	for i := 0; i < len(locations); i++ {
		fmt.Printf("%d: (%d %d)\n", i, locations[i].X, locations[i].Y)
	}
	var input string
	for {
		fmt.Print("Choice: ")
		fmt.Scanln(&input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		if choice < 0 || choice >= len(locations) {
			continue
		}
		return choice
	}
}
