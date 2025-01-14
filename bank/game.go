package bank

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"slices"
	"strings"
)

const (
	NumRounds = 20
)

func NewGame(players ...PlayerStrategy) *Game {
	return newGameWithDie(SixSidedDie{}, players...)
}

func newGameWithDie(dice Dice, players ...PlayerStrategy) *Game {
	result := &Game{dice: dice}
	for _, player := range players {
		result.Players = append(result.Players, NewPlayer(player, result))
	}
	return result
}

type Game struct {
	Players []*Player
	Round   Round
	dice    Dice
}

func (g *Game) Play() {
	g.playWithBuffer(os.Stdout)
}

func (g *Game) playWithBuffer(writer io.Writer) {
	for i := 0; i < NumRounds; i++ {
		g.NewRound(i + 1)
		for g.Round.Roll() && g.GetNumActivePlayers() > 0 {
			g.HandlePlayers()
		}
		g.ResetPlayers()
		fmt.Fprintln(writer, g.String())
	}
}

func (g *Game) NewRound(roundNumber int) {
	g.Round = NewRound(roundNumber, g.dice)
}

func (g *Game) HandlePlayers() {
	for g.handlePlayersOnce() {
	}
}

func (g *Game) handlePlayersOnce() bool {
	newBankedPlayers := g.PlayersPlay(false)
	if newBankedPlayers > 0 { // if somebody banked, restart the loop
		return true
	}
	newBankedPlayers = g.PlayersPlay(true)
	if newBankedPlayers == 0 { // if nobody banked, exit the loop
		return false
	}
	return true
}

func (g *Game) PlayersPlay(lastChance bool) (newBankedPlayers int) {
	beforeCount := g.GetNumBankedPlayers()
	for _, player := range g.shufflePlayers() {
		gameSnapShot := g.CopyGameInfo()
		if lastChance {
			player.LastChance(gameSnapShot, player.PlayerInfo)
		} else {
			player.Play(gameSnapShot, player.PlayerInfo)
		}
	}
	afterCount := g.GetNumBankedPlayers()
	return afterCount - beforeCount
}

func (g *Game) shufflePlayers() (result []*Player) {
	copyList := slices.Clone(g.Players)
	for len(copyList) > 0 {
		randIndex := rand.Intn(len(copyList))
		result = append(result, copyList[randIndex])
		copyList = append(copyList[:randIndex], copyList[randIndex+1:]...)
	}
	return result
}

func (g *Game) ResetPlayers() {
	for _, player := range g.Players {
		player.RoundOver()
	}
}

func (g *Game) GetNumBankedPlayers() (result int) {
	for _, player := range g.Players {
		if player.IsBanked {
			result++
		}
	}
	return result
}

func (g *Game) GetNumActivePlayers() (result int) {
	for _, player := range g.Players {
		if !player.IsBanked {
			result++
		}
	}
	return result
}

func (g *Game) String() string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "Round %d\n", g.Round.RoundNumber)
	for i, player := range g.Players {
		fmt.Fprintf(sb, g.getPlayerFormatString(), i, player.strategy.GetName(), player.Score)
	}
	return sb.String()
}

func (g *Game) getPlayerFormatString() string {
	return fmt.Sprintf("Player %%d %%-%ds: %%d\n", g.getLongestName()+1)
}

func (g *Game) getLongestName() (longestNameLength int) {
	for _, player := range g.Players {
		if currNameLength := len(player.strategy.GetName()); currNameLength > longestNameLength {
			longestNameLength = currNameLength
		}
	}
	return longestNameLength
}

func (g *Game) CopyGameInfo() GameInfo {
	var playersCopy []PlayerInfo
	for _, player := range g.Players {
		playersCopy = append(playersCopy, player.PlayerInfo)
	}
	return GameInfo{
		Players: playersCopy,
		Round:   g.Round,
	}
}
