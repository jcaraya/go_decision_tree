package dataset

import (
	"math/rand"
	"time"
)

// makeShuffledRange creates a range from zero to max-1
func makeShuffledRange(max int) []int {
	r := MakeRange(max)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
	return r
}

// ShuffleDataset receives the (x, y) pair components of a dataset and returns a
// shuffled version of it.
func ShuffleDataset(x [][]float64, y []int) ([][]float64, []int) {
	if len(x) != len(y) {
		panic("provided components (x, y) have differnth length values")
	}

	// Create a copy of the (x, y) pair to avoid modifying the original data.
	newX := make([][]float64, len(x))
	newY := make([]int, len(y))

	r := makeShuffledRange(len(x))
	for current, shuffled := range r {
		newX[current] = x[shuffled]
		newY[current] = y[shuffled]
	}

	return newX, newY
}
