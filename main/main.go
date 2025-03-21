package main

import (
	"fmt"

	"bank/bank"
	"bank/easton"
	"bank/example"
)

func main() {
	PlayGame()
}

func PlayGame() {
	game := bank.NewGame(
		&example.ExamplePlayer{RoundLimit: 20},
		&example.ExamplePlayer{RoundLimit: 100},
		&example.ExamplePlayer{RoundLimit: 200},
		&example.AnotherExamplePlayer{},
		&example.BankAfter{BankAfter: 5},
		&example.BankAfter{BankAfter: 6},
		&easton.StrategicDoubler{},
		&easton.Easton{},
	)
	//game.ActualBusiness()
	game.Play()
	game.ResetPlayers()
}

func roundExperiment() {
	bestRound := bank.Round{}
	for true {
		r := bank.Round{}
		r.Play()
		if r.Score > bestRound.Score {
			bestRound = r
			fmt.Println(bestRound.String())
		}
	}
}
