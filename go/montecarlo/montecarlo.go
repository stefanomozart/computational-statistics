// Package montecarlo implements a few montecarlo simullations
package montecarlo

import (
	"math/big"
	"time"

	"github.com/stefanomozart/paillier"
)

// Simullation is a runnable
type Simullation interface {
	Setup(args ...interface{})
	Run() error
	TimeElapsed() int64
}

// Run `sim` (the given Simullation function) for `m` times
func Run(sim Simullation, m int, args ...interface{}) {
	sim.Setup(args)
	times := make([]int64, m)
	for i := 0; i < m; i++ {
		sim.Run()
		times[i] = sim.TimeElapsed()
	}

}

// MeanWithHomomorphic is the computation of a mean
type MeanWithHomomorphic struct {
	numbers  []int64
	mean     int64
	duration time.Duration
}

// Setup the data needed for simullation
func (ms *MeanWithHomomorphic) Setup(args ...interface{}) {
	// unpack the generic args into numbers []int64
	ms.numbers = append(numbers, args...)
}

// Run the mean calculation simmulation
func (ms *MeanWithHomomorphic) Run() error {
	start := time.Now()
	pk, sk, err := paillier.GenerateKeyPair(2048)
	if err != nil {
		return err
	}

	cts := make([]*big.Int, len(ms.numbers))
	for i, m := range ms.numbers {
		cts[i], err = pk.Encrypt(m)
	}
	ctSum := pk.BatchAdd(cts...)
	ctMean, err := pk.DivPlaintext(ctSum, int64(len(ms.numbers)))
	ms.mean, err = sk.Decrypt(ctMean)
	if err != nil {
		return err
	}

	ms.duration = time.Since(start)
	return nil
}

// TimeElapsed during simullation
func (ms *MeanWithHomomorphic) TimeElapsed() int64 {
	return ms.duration.Nanoseconds()
}
