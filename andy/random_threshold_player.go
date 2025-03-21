package andy

import (
	"fmt"
	"math/rand"

	"bank/bank"
)

type RandomThresholdPlayer struct {
	bank.PlayerControls
	RoundLimit     int
	threshold      float64
	currentRound   int
	roundChange    bool
	previousScore  int
	currentCompare int
}

func (this RandomThresholdPlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {}

func (this *RandomThresholdPlayer) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	this.determinRandomThreshold()
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

func (this RandomThresholdPlayer) GetName() string {
	return fmt.Sprintf("Andy's Random Threshold Player")
}

func (this *RandomThresholdPlayer) determinRandomThreshold() {
	random := rand.Intn(3)
	switch random {
	case 0:
		this.threshold -= .1
	case 1:
		this.threshold += .1
	case 2:
		//do nothing
	}
	if this.threshold >= 0 || this.threshold > 1 {
		this.threshold = .4
	}
}
