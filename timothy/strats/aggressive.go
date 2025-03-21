package strats

import (
	"bank/bank"
)

type Aggressive struct {
	banked bool
	target int
}

// if nobody banks
func (this *Aggressive) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) (bank bool) {
	return false
}

// called each round or each time a player banks
func (this *Aggressive) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) (bank bool) {
	if this.target == 0 {
		this.target = 201
	}

	if game.Round.Score >= this.target {
		this.banked = true
		return true
	}

	return false
}

func (this *Aggressive) Reset() {
	if this.banked {
		this.target += 1
		this.banked = false
		return
	}

	this.target -= 1
}

func (this *Aggressive) Name() string {
	return "Aggressive"
}
