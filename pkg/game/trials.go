package game

import (
	"context"
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
	ctx := context.Background()
	results := make(chan int)
	for i := 0; i < 10; i++ {
		go runTrialsWorker(ctx, p, results)
	}
	finished := false
	for {
		if finished {
			break
		}
		select {
		case score := <-results:

			trials = append(trials, float64(score))
		default:
			if isFinished(len(trials)) {
				ctx.Done()
				finished = true
			}
		}
		//log.Infof("Finished game, score: %d", score)
	}
	fmt.Printf("Ran %d trials of %v in %s.\n", len(trials), reflect.TypeOf(p), time.Now().Sub(start))
	fmt.Printf("Mean: %f, standard deviation: %f\n", stats.Mean(trials), stats.StdDev(trials))
}

func runTrialsWorker(ctx context.Context, p Player, results chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			g := New()
			score, _ := g.Play(p)
			results <- score
		}
	}
}

func RunTrialsByTime(p Player, duration time.Duration) {
	finish := time.Now().Add(duration)
	runTrials(p, func(trial int) bool {
		return time.Now().After(finish)
	})
}
