package carter

import (
	"bank/bank"
)

type Carter struct {
	bank.PlayerControls
	bankNextRoll bool
}

func (this *Carter) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {

	myScore := yourInfo.Score
	winning := true
	for _, player := range game.Players {
		if game.Round.CurrentDieRoll+myScore <= player.Score {
			winning = false
		}
	}
	//Hail Mary
	if game.Round.RoundNumber == 20 && !winning {
		//Dont bank!
	} else {
		if winning && game.Round.Score > 100 {
			this.Bank()
		}

		if this.bankNextRoll {
			this.Bank()
		}
		if game.GetNumBankedPlayers() == len(game.Players)-1 {
			this.bankNextRoll = true
		}

		if game.Round.Score > 300 {
			this.Bank()
		}
	}

}

// LastChance is called if nobody banked. It is not necessary to use this function but it could play into your
// strategy
func (this *Carter) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {

}

// RoundOver is called when the round ends. You can use this method to reset any internal data you store on your
// player
func (this *Carter) RoundOver() {
	this.bankNextRoll = false
}

// GetName used for display
func (this *Carter) GetName() string {
	return "Carter"
}
