package timothy

import (
	"bank/bank"
	"bank/timothy/strats"
	"fmt"
)

type Bot struct {
	bank.PlayerControls
	gameHistory []winCount
	roundInfo   bank.Round
	players     []bank.PlayerInfo
	strategy    strats.Strategy
	gameNumber  int
}

func (this *Bot) GetName() string {
	return "Timothy"
}

// if nobody banks
func (this *Bot) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	this.players = game.Players
	if this.strategy.LastChance(game, yourInfo) {
		this.Bank()
	}
}

// called each round or each time a player banks
func (this *Bot) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	this.players = game.Players
	this.initialize(game.Players)
	this.roundInfo = game.Round
	if this.strategy.Play(game, yourInfo) {
		this.Bank()
	}
}

func (this *Bot) RoundOver() {
	this.strategy.Reset()
	if this.roundInfo.RoundNumber >= bank.NumRounds {
		this.gameOver()
	}
}

func (this *Bot) PrintTotalWins() {
	for _, wins := range this.gameHistory {
		fmt.Printf("%s: %d\n", wins.Name, wins.Count)
	}
}

func (this *Bot) initialize(players []bank.PlayerInfo) {
	if len(this.gameHistory) == 0 {
		this.strategy = new(strats.Aggressive)
		for _, player := range players {
			this.gameHistory = append(this.gameHistory, winCount{Count: 0, Name: player.Name, ID: player.ID})
		}
	}
}

func (this *Bot) gameOver() {
	this.gameNumber++
	best := this.players[0]
	for _, current := range this.players {
		if current.Score > best.Score {
			best = current
		}
	}

	this.incrementWins(best.Score)
	if this.gameNumber%50 == 0 && this.gameNumber > 0 {
		this.reevaluateStrat(best.Score, best.ID)
		fmt.Printf("%s\n", this.strategy.Name())
	}
}

func (this *Bot) incrementWins(score int) {
	for _, player := range this.players {
		if player.Score == score {
			this.incrementWin(player.ID)
		}
	}
}

func (this *Bot) incrementWin(id int) {
	for iHistory := range this.gameHistory {
		if this.gameHistory[iHistory].ID == id {
			this.gameHistory[iHistory].Count++
		}
	}
}

func (this *Bot) reevaluateStrat(bestScore int, bestID int) {
	fmt.Printf("reevaluating -> ")
	if this.inTheLead(bestScore) {
		if _, ok := this.strategy.(*strats.Aggressive); ok {
			return
		}

		this.strategy = new(strats.Aggressive)
		return
	}

	this.strategy = &strats.Chase{ID: bestID}
}

func (this *Bot) inTheLead(bestScore int) bool {
	for _, player := range this.players {
		if player.Score >= bestScore-10 {
			if player.Name == this.GetName() {
				return true
			}
		}
	}

	return false
}
