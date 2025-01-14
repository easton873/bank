package bank

import (
	"fmt"
	"strings"
	"testing"
)

func assert(t *testing.T, test bool) {
	t.Helper()
	if !test {
		t.Errorf("assert failed")
	}
}

func TestPlayersBanking(t *testing.T) {
	p1 := &FakeBankingPlayer{}
	game := NewGame(p1)
	game.Round.Score = 10

	game.handlePlayersOnce()
	assert(t, game.Players[0].Score == 10)
}

func TestUnbankingPlayers(t *testing.T) {
	p1 := &FakePlayer{}
	p1.shouldBank = true
	game := NewGame(p1)

	assert(t, !game.Players[0].IsBanked)
	game.handlePlayersOnce()
	assert(t, game.Players[0].IsBanked)
	game.ResetPlayers()
	assert(t, !game.Players[0].IsBanked)
}

func TestPlayersPlayCorrectlyWhenOthersBank(t *testing.T) {
	p1 := &FakePlayer{}
	p2 := &FakePlayer{}
	p3 := &FakePlayer{}
	game := NewGame(
		p1,
		p2,
		p3,
	)
	result := game.handlePlayersOnce()
	assert(t, !result)
	assert(t, p1.playCount == 1)
	assert(t, p2.playCount == 1)
	assert(t, p3.playCount == 1)
	assert(t, p1.lastChanceCount == 1)
	assert(t, p2.lastChanceCount == 1)
	assert(t, p3.lastChanceCount == 1)

	p1.shouldBank = true
	result = game.handlePlayersOnce()
	assert(t, result)
	result = game.handlePlayersOnce()
	assert(t, !result)
	assert(t, p1.playCount == 2)
	assert(t, p2.playCount == 3)
	assert(t, p3.playCount == 3)
	assert(t, p1.lastChanceCount == 1)
	assert(t, p2.lastChanceCount == 2)
	assert(t, p3.lastChanceCount == 2)

	p2.shouldLastChanceBank = true
	result = game.handlePlayersOnce()
	assert(t, result)
	assert(t, p1.playCount == 2)
	assert(t, p2.playCount == 4)
	assert(t, p3.playCount == 4)
	assert(t, p1.lastChanceCount == 1)
	assert(t, p2.lastChanceCount == 3)
	assert(t, p3.lastChanceCount == 3)
	result = game.handlePlayersOnce()
	assert(t, !result)
	assert(t, p1.playCount == 2)
	assert(t, p2.playCount == 4)
	assert(t, p3.playCount == 5)
	assert(t, p1.lastChanceCount == 1)
	assert(t, p2.lastChanceCount == 3)
	assert(t, p3.lastChanceCount == 4)
}

func Test2Rounds(t *testing.T) {
	p1 := &FakePlayer{}
	p2 := &FakePlayer{}
	p3 := &FakePlayer{}

	game := newGameWithDie(&FixedDice{}, p1, p2, p3)
	game.NewRound(1)

	// Roll 1
	game.Round.Roll()
	assert(t, game.Round.CurrentDieRoll == 1)

	// Roll 2
	p1.shouldBank = true
	game.Round.Roll()
	res := game.handlePlayersOnce()
	assert(t, res)
	assert(t, game.Players[0].Score == 2)

	// Roll 3
	p2.shouldBank = true
	game.Round.Roll()
	res = game.handlePlayersOnce()
	assert(t, res)
	assert(t, game.Players[1].Score == 5)

	// Roll 4
	p3.shouldBank = true
	game.Round.Roll()
	res = game.handlePlayersOnce()
	assert(t, res)
	assert(t, game.Players[2].Score == 9)
	res = game.handlePlayersOnce()
	assert(t, !res)

	// Roll 5
	game.NewRound(2)
	game.ResetPlayers()
	game.Round.Roll()
	res = game.handlePlayersOnce()
	assert(t, res)
	res = game.handlePlayersOnce()
	assert(t, !res)
	assert(t, game.Players[0].Score == 7)
	assert(t, game.Players[1].Score == 10)
	assert(t, game.Players[2].Score == 14)
}

func TestFullGame(t *testing.T) {
	p1 := &BankAfterPlayer{BankAfter: 1}
	p2 := &BankAfterPlayer{BankAfter: 2}
	p3 := &BankAfterPlayer{BankAfter: 3}

	game := newGameWithDie(&FixedDice{}, p1, p2, p3)
	game.playWithBuffer(&strings.Builder{})

	assert(t, game.Players[0].Score == 78)
	assert(t, game.Players[1].Score == 58)
	assert(t, game.Players[2].Score == 86)
}

// mocks

type FakeBankingPlayer struct {
	PlayerControls
}

func (f FakeBankingPlayer) Play(game GameInfo, yourInfo PlayerInfo) {
	f.Bank()
	f.Bank()
	f.Bank()
	f.Bank()
	f.Bank()
	f.Bank()
	f.Bank()
	f.Bank()
}

func (f FakeBankingPlayer) LastChance(game GameInfo, yourInfo PlayerInfo) {
	f.Bank()
	f.Bank()
	f.Bank()
}

func (f FakeBankingPlayer) GetName() string {
	return "FakeBankingPlayer"
}

type FakePlayer struct {
	PlayerControls
	playCount            int
	lastChanceCount      int
	shouldBank           bool
	shouldLastChanceBank bool
}

func (f *FakePlayer) Play(game GameInfo, yourInfo PlayerInfo) {
	if f.shouldBank {
		f.Bank()
	}
	f.playCount++
}

func (f *FakePlayer) LastChance(game GameInfo, yourInfo PlayerInfo) {
	if f.shouldLastChanceBank {
		f.Bank()
	}
	f.lastChanceCount++
}

func (f *FakePlayer) GetName() string {
	return "FakePlayer"
}

type BankAfterPlayer struct {
	PlayerControls
	BankAfter int
}

func (b BankAfterPlayer) Play(game GameInfo, yourInfo PlayerInfo) {
	if game.Round.NumRolls == b.BankAfter {
		b.Bank()
	}
}

func (b BankAfterPlayer) GetName() string {
	return fmt.Sprintf("Bank After %d Rolls", b.BankAfter)
}

type FixedDice struct {
	lastRoll int
}

func (f *FixedDice) Roll() int {
	f.lastRoll++
	if f.lastRoll == 7 {
		f.lastRoll = 1
	}
	return f.lastRoll
}
