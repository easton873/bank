package andy

import (
	"fmt"
	"math/rand"

	"bank/bank"
)

type RandomPlayer struct {
	bank.PlayerControls
}

func (this RandomPlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if rand.Intn(2) == 0 {
		this.Bank()
	}
}

func (this RandomPlayer) GetName() string {
	return fmt.Sprintf("Andy's Random Player")
}
