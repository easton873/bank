package ryan

import "bank/bank"

type RyansAllPowerfulBot struct {
	bank.PlayerControls
	nextRoll bool
}

func (a *RyansAllPowerfulBot) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.Round.RoundNumber >= game.Round.NumRolls {
		a.Bank()
		a.nextRoll = false
	}
}

func (a *RyansAllPowerfulBot) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.GetNumNOTBankedPlayers() == 5 {
		a.nextRoll = true
	}
}

func (a *RyansAllPowerfulBot) RoundOver() {
	a.nextRoll = false
}

func (a *RyansAllPowerfulBot) GetName() string {
	return "I have no idea what I'm doing."
}
