package main

import (
	"github.com/lukemassa/onezero/pkg/game"
	"github.com/lukemassa/onezero/pkg/game/players"
	log "github.com/sirupsen/logrus"
)

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	//player := players.PreferClearingPlayer{}
	//	player := players.UpperLeftPlayer{}
	// player := players.TextBasedPlayer{}
	player := players.TUIPlayer{}
	//game.RunTrials(&player, 100_00)
	//game.RunTrialsByTime(player, 10*time.Second)
	game.RunInteractive(player)

}
