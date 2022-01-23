package game

import "fmt"

//import log "github.com/sirupsen/logrus"

func RunTrials(p Player, numTrials int) {
	totalScore := 0.0
	for i := 0; i < numTrials; i++ {
		//log.Info("Starting game...")
		g := New(p)
		score, _ := g.Play()
		totalScore += float64(score)
		//log.Infof("Finished game, score: %d", score)
	}
	fmt.Printf("Average score: %f\n", totalScore/float64(numTrials))
}
