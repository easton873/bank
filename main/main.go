package main

import (
	"fmt"

	"bank/aubry"
	"bank/bank"
	"bank/example"
)

func main() {
	//PlayGame()
	PlayManyGames(10_000)
}

func PlayGame() {
	game := bank.NewGame(
		&example.ExamplePlayer{RoundLimit: 20},
		&example.ExamplePlayer{RoundLimit: 100},
		&example.ExamplePlayer{RoundLimit: 200},
		&example.AnotherExamplePlayer{},
		&example.BankAfter{BankAfter: 5},
		&example.BankAfter{BankAfter: 6},
	)
	game.Play()
}

func PlayManyGames(rounds int) float64 {
	p1wins := 0.0
	p2wins := 0.0
	p3wins := 0.0
	p4wins := 0.0
	p5wins := 0.0
	p6wins := 0.0
	p7wins := 0.0

	for i := 0; i < rounds; i++ {
		game := bank.NewGame(
			&example.ExamplePlayer{RoundLimit: 20},
			&example.ExamplePlayer{RoundLimit: 100},
			&example.ExamplePlayer{RoundLimit: 200},
			&example.AnotherExamplePlayer{},
			&example.BankAfter{BankAfter: 5},
			&example.BankAfter{BankAfter: 6},
			&aubry.Aubry{},
		)
		game.Play()
		numPlayers := len(game.Players)
		maxScore := -1
		maxPlayer := ""
		for i := 0; i < numPlayers; i++ {
			if game.Players[i].Score > maxScore {
				maxScore = game.Players[i].Score
				maxPlayer = game.Players[i].Name
			}
		}
		if maxPlayer == "Bank Once At 20" {
			p1wins++
		}
		if maxPlayer == "Bank Once At 100" {
			p2wins++
		}
		if maxPlayer == "Bank Once At 200" {
			p3wins++
		}
		if maxPlayer == "Bank 1 After Everybody" {
			p4wins++
		}
		if maxPlayer == "Bank After 5 Rolls" {
			p5wins++
		}
		if maxPlayer == "Bank After 6 Rolls" {
			p6wins++
		}
		if maxPlayer == "Aubry's Special Bot" {
			p7wins++
		}
	}
	fmt.Printf("p1 wins: %f\n", p1wins/float64(rounds)*100)
	fmt.Printf("p2 wins: %f\n", p2wins/float64(rounds)*100)
	fmt.Printf("p3 wins: %f\n", p3wins/float64(rounds)*100)
	fmt.Printf("p4 wins: %f\n", p4wins/float64(rounds)*100)
	fmt.Printf("p5 wins: %f\n", p5wins/float64(rounds)*100)
	fmt.Printf("p6 wins: %f\n", p6wins/float64(rounds)*100)
	fmt.Printf("Aubry wins: %f\n", p7wins/float64(rounds)*100)
	return p7wins / float64(rounds)
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
