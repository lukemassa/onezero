package main

import (
	"time"

	"github.com/lukemassa/onezero/pkg/game"
	"github.com/lukemassa/onezero/pkg/game/players"
	log "github.com/sirupsen/logrus"
)

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	player := players.NewRandomPlayer()
	//game.RunTrials(&player, 10)
	game.RunTrialsByTime(&player, 10*time.Second)
	// g := game.New(&player)
	// score, lastPiece := g.Play()
	// fmt.Printf("Game over! Score: %d\n", score)
	// g.Show()
	// fmt.Printf("Could not place piece:\n%s\n", lastPiece.String())
}
