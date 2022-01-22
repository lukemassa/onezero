package main

import (
	"fmt"

	"github.com/lukemassa/onezero/pkg/game"
	"github.com/lukemassa/onezero/pkg/game/players"
	log "github.com/sirupsen/logrus"
)

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	player := players.InteractivePlayer{}
	g := game.New(player)
	score, lastPiece := g.Play()
	fmt.Printf("Game over! Score: %d\n", score)
	g.Show()
	fmt.Printf("Could not place piece:\n%s\n", lastPiece.String())
}
