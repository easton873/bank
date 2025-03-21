package kent

import (
	"fmt"

	"bank/bank"
)

// Player make sure to embed bank.PlayerControls into your player or you won't be able to implement the
// bank.PlayerStrategy interface
type Player struct {
	bank.PlayerControls
	RoundLimit int
}

func (this Player) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	//if game.Round.Score >= this.RoundLimit || game.Round.NumRolls == 500 {
	//randomNumber := rand.Intn(50) + 1

	if game.Round.NumRolls > 5 {
		//fmt.Println(randomNumber)
		//if game.Round.NumRolls > randomNumber {
		this.Bank()
	}
}

func (this Player) GetName() string {
	return fmt.Sprintf("Bank Once At %d", this.RoundLimit)
}
