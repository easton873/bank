package xan

import (
	"fmt"

	"bank/bank"
)

const roundLimit = 10

// ExamplePlayer make sure to embed bank.PlayerControls into your player or you won't be able to implement the
// bank.PlayerStrategy interface
type XanPlayer struct {
	bank.PlayerControls
}

func (e XanPlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.Round.RoundNumber > roundLimit {
		e.Bank()
	}
}

func (e XanPlayer) GetName() string {
	return fmt.Sprintf("XanPlayer")
}
