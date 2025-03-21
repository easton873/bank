package savannah

import (
	"fmt"

	"bank/bank"
)

type SavannahPlayer struct {
	bank.PlayerControls
	nextRoll bool
}

func (this *SavannahPlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	response := GetMagic8BallResponse(game.Round.RoundNumber)
	fmt.Println("🎱 Magic 8 Ball says:", response)

	this.LastChance(game, yourInfo)

	switch response {
	case ItIsCertain, YesDefinitely:
		this.Bank()           // Choose to bank the points
		this.nextRoll = false // End turn
	default:
		this.nextRoll = true // Roll again
	}
}

func (this *SavannahPlayer) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.GetNumNOTBankedPlayers() == 1 {
		this.Bank()           // Choose to bank the points
		this.nextRoll = false // End turn
	}
}

func (this *SavannahPlayer) RoundOver() {
	this.nextRoll = false
}

func (this *SavannahPlayer) GetName() string {
	return "Bank 1 After Everybody"
}
