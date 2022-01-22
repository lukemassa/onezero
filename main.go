package main

import (
	"fmt"

	"github.com/lukemassa/onezero/pkg/game"
)

func main() {

	player := game.InteractivePlayer{}
	g := game.New(player)
	score, lastPiece := g.Play()
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Game over! Score: %d\n", score)
	g.Show()
	fmt.Printf("Could not place piece:\n%s\n", lastPiece.String())
}
