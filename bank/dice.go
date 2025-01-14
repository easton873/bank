package bank

import "math/rand"

type Dice interface {
	Roll() int
}

type SixSidedDie struct{}

func (s SixSidedDie) Roll() int {
	return randInt(1, 6)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
