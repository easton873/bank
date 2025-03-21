package adam

import (
	"bank/bank"
)

// ExamplePlayer make sure to embed bank.PlayerControls into your player or you won't be able to implement the
// bank.PlayerStrategy interface
type AdamWinByTwo struct {
	bank.PlayerControls
	TurnsTo        int
	leaderScore    int
	finalCountdown bool
	currentDieRoll int
}

func (this *AdamWinByTwo) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {

	currentLeaderIndex := -1
	leaderScore := 0

	if this.finalCountdown {
		return
	}

	for i, player := range game.Players {

		if player.Name == yourInfo.Name {
			continue
		}

		if player.Score > leaderScore {
			leaderScore = player.Score
			currentLeaderIndex = i
		}
	}

	if currentLeaderIndex != -1 {
		if game.Players[currentLeaderIndex].IsBanked {
			this.TurnsTo = 2
			this.finalCountdown = true
		}
	}

}

func (this *AdamWinByTwo) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if this.currentDieRoll != game.Round.CurrentDieRoll {
		this.TurnsTo--
		this.currentDieRoll = game.Round.CurrentDieRoll
	}
	if !this.finalCountdown {
		if game.Round.Score > 500 {
			this.Bank()
		}
	}
	if this.TurnsTo == 0 {
		this.Bank()
	}
}

func (this *AdamWinByTwo) GetName() string {
	return "AdamUltraBot"
}

func (this *AdamWinByTwo) RoundOver() {
	this.TurnsTo = 15
	this.currentDieRoll = -1
}
