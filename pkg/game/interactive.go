package game

import (
	"fmt"
	"strconv"
)

type InteractivePlayer struct{}

func (i InteractivePlayer) Decide(b Board, p Piece, locations []Location) int {
	fmt.Print("\033[H\033[2J")
	b.Show()
	p.Show()
	fmt.Println("Options:")
	for i := 0; i < len(locations); i++ {
		fmt.Printf("%d: (%d %d)\n", i, locations[i].x, locations[i].y)
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
