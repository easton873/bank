package bank

type GameInfo struct {
	Players []PlayerInfo
	Round   Round
}

func (g *GameInfo) GetNumBankedPlayers() (result int) {
	for _, player := range g.Players {
		if player.IsBanked {
			result++
		}
	}
	return result
}

func (g *GameInfo) GetNumNOTBankedPlayers() (result int) {
	for _, player := range g.Players {
		if !player.IsBanked {
			result++
		}
	}
	return result
}
