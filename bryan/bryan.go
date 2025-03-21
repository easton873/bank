package bryan

import (
	"bank/bank"
)

const defaultRoundLimit = 395

type BryanPlayer struct {
	bank.PlayerControls
	RoundLimit int
}

func (this *BryanPlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if this.RoundLimit == 0 {
		this.RoundLimit = defaultRoundLimit
	}
	
	if game.Round.Score >= this.RoundLimit {
		this.Bank()
	}
}

func (this *BryanPlayer) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {}

func (this *BryanPlayer) RoundOver() {}

func (this *BryanPlayer) GetName() string {
	return "Bryan"
}
