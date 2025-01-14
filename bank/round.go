package bank

import "fmt"

type Round struct {
	dice           Dice
	RoundNumber    int
	NumRolls       int
	Score          int
	CurrentDieRoll int
}

func NewRound(roundNumber int, dice Dice) Round {
	return Round{
		dice:        dice,
		RoundNumber: roundNumber,
	}
}

func (r *Round) Play() {
	for r.Roll() {
	}
}

func (r *Round) Roll() bool {
	roll := r.dice.Roll()
	r.NumRolls++
	r.CurrentDieRoll = roll
	if r.NumRolls == 1 { // if first roll in round
		r.addFaceValue()
		return true
	} else if roll == 1 {
		r.Score = 0
		return false
	} else if roll == 2 {
		r.Score *= 2
		return true
	}
	r.addFaceValue()
	return true
}

func (r *Round) addFaceValue() {
	r.Score += r.CurrentDieRoll
}

func (r *Round) String() string {
	return fmt.Sprintf("Score: %d, NumRolls: %d", r.Score, r.NumRolls)
}
