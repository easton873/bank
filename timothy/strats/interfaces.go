package strats

import "bank/bank"

type Strategy interface {
	// if nobody banks
	LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) (bank bool)

	// called each round or each time a player banks
	Play(game bank.GameInfo, yourInfo bank.PlayerInfo) (bank bool)

	Reset()

	Name() string
}
