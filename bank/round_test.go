package bank

import "testing"

func TestRoll(t *testing.T) {
	dice := FakeDice{numToReturn: 1}
	round := Round{dice: &dice}

	// rolls one round ends
	res := round.Roll()
	assert(t, res)
	assert(t, round.CurrentDieRoll == 1)
	assert(t, round.Score == 1)
	res = round.Roll()
	assert(t, !res)
	assert(t, round.Score == 0)

	// score resets
	dice.numToReturn = 5
	res = round.Roll()
	assert(t, res)
	assert(t, round.CurrentDieRoll == 5)
	assert(t, round.Score == 5)
	dice.numToReturn = 1
	res = round.Roll()
	assert(t, !res)
	assert(t, round.CurrentDieRoll == 1)
	assert(t, round.Score == 0)

	// score doubles
	dice.numToReturn = 5
	round.Roll()
	res = round.Roll()
	assert(t, res)
	assert(t, round.Score == 10)
	dice.numToReturn = 2
	res = round.Roll()
	assert(t, res)
	assert(t, round.Score == 20)
}

type FakeDice struct {
	numToReturn int
}

func (f *FakeDice) Roll() int {
	return f.numToReturn
}
