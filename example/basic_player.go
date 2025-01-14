package example

import (
	"fmt"

	"bank/bank"
)

// ExamplePlayer make sure to embed bank.PlayerControls into your player or you won't be able to implement the
// bank.PlayerStrategy interface
type ExamplePlayer struct {
	bank.PlayerControls
	RoundLimit int
}

func (e ExamplePlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.Round.Score >= e.RoundLimit {
		e.Bank()
	}
}

func (e ExamplePlayer) GetName() string {
	return fmt.Sprintf("Bank Once At %d", e.RoundLimit)
}
