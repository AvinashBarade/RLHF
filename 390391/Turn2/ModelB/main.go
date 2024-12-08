package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// Beta struct with methods for sampling
type Beta struct {
	alpha, beta float64
}

// NewBeta creates a new Beta distribution
func NewBeta(alpha, beta int) *Beta {
	return &Beta{float64(alpha), float64(beta)}
}

// Update updates the alpha and beta for successes and failures
func (b *Beta) Update(successes, failures int) {
	b.alpha += float64(successes)
	b.beta += float64(failures)
}

// Sample returns a sample from the beta distribution
func (b *Beta) Sample(rnd *rand.Rand) float64 {
	betaDist := distuv.Beta{Alpha: b.alpha, Beta: b.beta, Src: rnd}
	return betaDist.Rand()
}

// CompareABTest calculates the probability that B > A
func CompareABTest(priorA, priorB *Beta, numSamples int) float64 {
	rnd := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	bWins := 0

	// Monte Carlo simulation to compare A and B
	for i := 0; i < numSamples; i++ {
		sampleA := priorA.Sample(rnd)
		sampleB := priorB.Sample(rnd)
		if sampleB > sampleA {
			bWins++
		}
	}
	return float64(bWins) / float64(numSamples)
}

func main() {
	// Initial prior for both groups A and B
	priorA := NewBeta(1, 1)
	priorB := NewBeta(1, 1)

	// Simulated data for A/B test
	// Group A: 20 successes out of 100 trials
	// Group B: 30 successes out of 100 trials
	priorA.Update(20, 80)
	priorB.Update(30, 70)

	// Calculate probability that B beats A, using 100,000 samples for precision
	probBBeatsA := CompareABTest(priorA, priorB, 100000)
	fmt.Printf("Probability that Group B outperforms Group A: %.4f\n", probBBeatsA)
}
