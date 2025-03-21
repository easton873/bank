package thomas

import (
	"bank/bank"
)

type Thomas struct {
	bank.PlayerControls
	bryanBanked bool
}

func (t *Thomas) otherPlayerInfo(players []bank.PlayerInfo, yourInfo bank.PlayerInfo) []bank.PlayerInfo {
	for i, player := range players {
		if player == yourInfo {
			return append(players[:i], players[i+1:]...)
		}
	}
	return nil
}

func (t *Thomas) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	t.normalStrategy(game, yourInfo)
}

func (t *Thomas) normalStrategy(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if t.bryanBanked {
		t.Bank()
	}
}

func (t *Thomas) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	t.Play(game, yourInfo)
	for _, player := range game.Players {
		if player.Name == "Bryan" {
			if player.IsBanked {
				t.bryanBanked = true
			}
		}
	}
}

func (t *Thomas) RoundOver() {
	t.bryanBanked = false
}

func (t *Thomas) GetName() string {
	return "Thomas"
}
