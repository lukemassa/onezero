package players

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lukemassa/onezero/pkg/game"
)

type TextBasedPlayer struct{}

func (t TextBasedPlayer) Decide(b game.Board, p game.Piece, locations []game.Location) int {
	fmt.Print("\033[H\033[2J")
	b.Show()
	p.Show()
	fmt.Println("Options:")
	for i := 0; i < len(locations); i++ {
		//Make options and locations 1-indexed for human readability
		fmt.Printf("%d: (%d %d)\n", i+1, locations[i].X+1, locations[i].Y+1)
	}
	var input string
	for {
		fmt.Print("Choice: ")
		fmt.Scanln(&input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		//Choices are 1-indexed, locations are 0-indexed
		choice--
		if choice < 0 || choice >= len(locations) {
			continue
		}
		return choice
	}
}

func (i TextBasedPlayer) Move(b game.Board, p game.Piece) *game.Location {
	fmt.Print("\033[H\033[2J")
	size := len(b)
	b.Show()
	p.Show()
	fmt.Println("Options:")
	possibleLocations := b.GetPossibleLocations(p)
	if len(possibleLocations) == 0 {
		return nil
	}
	var input string
	for {
		fmt.Print("Choice: ")
		fmt.Scanln(&input)
		coordinates := strings.Split(input, ",")
		if len(coordinates) != 2 {
			continue
		}
		x_coordinate, err := strconv.Atoi(coordinates[0])
		if err != nil {
			continue
		}
		y_coordinate, err := strconv.Atoi(coordinates[1])
		if err != nil {
			continue
		}
		//Choices are 1-indexed, locations are 0-indexed
		x_coordinate--
		y_coordinate--
		if x_coordinate < 0 || x_coordinate >= size {
			continue
		}
		if y_coordinate < 0 || y_coordinate >= size {
			continue
		}
		ret := game.Location{X: x_coordinate, Y: y_coordinate}
		for i := 0; i < len(possibleLocations); i++ {
			if ret.X == possibleLocations[i].X && ret.Y == possibleLocations[i].Y {
				return &ret
			}
		}
		fmt.Println("Not a valid move, try again")
	}
}
