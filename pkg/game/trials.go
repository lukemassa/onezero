package game

import (
	"fmt"
	"reflect"
	"time"

	"github.com/aclements/go-moremath/stats"
)

//import log "github.com/sirupsen/logrus"

func RunTrials(p Player, numTrials int) {
	runTrials(p, func(trial int) bool {
		return trial == numTrials
	})

}

func runTrials(p Player, isFinished func(int) bool) {
	trials := make([]float64, 0)
	start := time.Now()
	for {
		//log.Info("Starting game...")
		g := New(p)
		score, _ := g.Play()
		trials = append(trials, float64(score))
		if isFinished(len(trials)) {
			break
		}
		//log.Infof("Finished game, score: %d", score)
	}
	fmt.Printf("Ran %d trials of %v in %s.\n", len(trials), reflect.TypeOf(p), time.Now().Sub(start))
	fmt.Printf("Mean: %f, standard deviation: %f\n", stats.Mean(trials), stats.StdDev(trials))
}

func RunTrialsByTime(p Player, duration time.Duration) {
	finish := time.Now().Add(duration)
	runTrials(p, func(trial int) bool {
		return time.Now().After(finish)
	})
}
