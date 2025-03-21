package andy

import (
	"fmt"

	"bank/bank"
)

type ThresholdPlayer struct {
	bank.PlayerControls
	RoundLimit     int
	threshold      float64
	currentRound   int
	roundChange    bool
	previousScore  int
	currentCompare int
}

func (this ThresholdPlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {}

func (this *ThresholdPlayer) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	this.threshold = .4
	totalBankedGreater := 0
	for _, player := range game.Players {
		if yourInfo.Score+game.Round.Score > player.Score && player.IsBanked {
			totalBankedGreater++
		}
	}
	if float64(totalBankedGreater)/float64(len(game.Players)) >= this.threshold {
		this.Bank()
		this.roundChange = true
		this.previousScore = yourInfo.Score
	}
}

func (this ThresholdPlayer) GetName() string {
	return fmt.Sprintf("Andy's Threshold Player")
}
