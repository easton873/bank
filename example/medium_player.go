package example

import (
	"fmt"

	"bank/bank"
)

type BankAfter struct {
	bank.PlayerControls
	BankAfter int
}

func (b BankAfter) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.Round.NumRolls == b.BankAfter {
		b.Bank()
	}
}

func (b BankAfter) GetName() string {
	return fmt.Sprintf("Bank After %d Rolls", b.BankAfter)
}
