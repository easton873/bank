package example

import "bank/bank"

type AnotherExamplePlayer struct {
	bank.PlayerControls
	nextRoll bool
}

func (a *AnotherExamplePlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if a.nextRoll {
		a.Bank()
		a.nextRoll = false
	}
}

func (a *AnotherExamplePlayer) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.GetNumNOTBankedPlayers() == 1 {
		a.nextRoll = true
	}
}

func (a *AnotherExamplePlayer) RoundOver() {
	a.nextRoll = false
}

func (a *AnotherExamplePlayer) GetName() string {
	return "Bank 1 After Everybody"
}
