package bank

type BankFn func()

type PlayerControls struct {
	Bank BankFn
}

func (pc *PlayerControls) LastChance(_ GameInfo, _ PlayerInfo) {}

func (pc *PlayerControls) RoundOver() {}

func (pc *PlayerControls) setPlayerControls(controls *PlayerControls) {
	*pc = *controls
}

// PlayerStrategy is the interface you will want to implement to compete. The only way to fully implement this interface
// is to embed the PlayerControls struct
type PlayerStrategy interface {
	// Play is run for every player that hasn't banked in this round yet. If any players bank it will run through Play()
	// for everybody who hasn't banked again. This continues everybody has either banked or chosen not to bank twice in a
	// row (the second chance is in the LastChance method)
	Play(game GameInfo, yourInfo PlayerInfo)

	// LastChance is called if nobody banked. It is not necessary to use this function but it could play into your
	// strategy
	LastChance(game GameInfo, yourInfo PlayerInfo)

	// RoundOver is called when the round ends. You can use this method to reset any internal data you store on your
	// player
	RoundOver()

	// GetName used for display
	GetName() string

	// setPlayerControls is a private method you can only implement by embedding a PlayerControls
	setPlayerControls(*PlayerControls)
}

type PlayerInfo struct {
	Score    int
	IsBanked bool
	Name     string
	ID       int
}

type Player struct {
	PlayerInfo
	strategy PlayerStrategy
}

func NewPlayer(player PlayerStrategy, game *Game, id int) *Player {
	basePlayer := &PlayerControls{}
	result := &Player{
		strategy: player,
	}
	bankFn := func() {
		if result.IsBanked {
			return
		}
		result.IsBanked = true
		result.Score += game.Round.Score
	}
	basePlayer.Bank = bankFn
	player.setPlayerControls(basePlayer)
	result.Name = player.GetName()
	result.ID = id
	return result
}

func (p *Player) Play(game GameInfo, yourInfo PlayerInfo) {
	if p.IsBanked {
		return
	}
	p.strategy.Play(game, yourInfo)
}

func (p *Player) LastChance(game GameInfo, yourInfo PlayerInfo) {
	if p.IsBanked {
		return
	}
	p.strategy.LastChance(game, yourInfo)
}

func (p *Player) RoundOver() {
	p.IsBanked = false
	p.strategy.RoundOver()
}
