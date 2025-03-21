package aubry

import "bank/bank"

type Aubry struct {
	bank.PlayerControls
}

func (this *Aubry) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.Round.Score >= 200 || game.Round.RoundNumber >= 18 || game.GetNumNOTBankedPlayers() == 1 {
		this.Bank()
	}
}

//func (this *Aubry) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
//}
//
//func (this *Aubry) RoundOver() {
//}

func (this *Aubry) GetName() string {
	return "Aubry's Special Bot"
}
