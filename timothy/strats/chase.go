package strats

import (
	"bank/bank"
	"fmt"
)

type Chase struct {
	ID           int
	bankNextTurn bool
}

// if nobody banks
func (this *Chase) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) (bank bool) {
	for _, player := range game.Players {
		if player.ID == this.ID {
			if player.IsBanked {
				this.bankNextTurn = true
			}
		}
	}

	return false
}

// called each round or each time a player banks
func (this *Chase) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) (bank bool) {
	return this.bankNextTurn
}

func (this *Chase) Reset() {
	this.bankNextTurn = false
}

func (this *Chase) Name() string {
	return fmt.Sprintf("Chase %d", this.ID)
}
