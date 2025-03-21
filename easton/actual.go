package easton

import "bank/bank"

type Easton struct {
	bank.PlayerControls
}

func (this Easton) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	winnerScore := getMaxScoreThatIsntMe(game, yourInfo)
	if winnerScore > yourInfo.Score && game.Round.Score+yourInfo.Score > winnerScore {
		this.Bank()
	}
	if game.GetNumNOTBankedPlayers() > 1 {
		return
	}
	this.Bank()
}

func (this Easton) GetName() string {
	return "Easton"
}

func getMaxScoreThatIsntMe(game bank.GameInfo, yourInfo bank.PlayerInfo) (result int) {
	for _, player := range game.Players {
		if player == yourInfo {
			continue
		}
		if player.Score > result {
			result = player.Score
		}
	}
	return result
}
