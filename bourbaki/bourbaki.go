package bourbaki

import (
	"math/rand"
	"os/exec"

	"bank/bank"
)

type NicolasPlayer struct {
	bank.PlayerControls
	nextRoll bool
}

func (this *NicolasPlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	messages := []string{"all our base is belong to us", "I'm so weak", "nico", "east is up", "I am clancy", "sudo woodo", "Easton, I'm not feeling well"}

	target := game.Players[rand.Intn(len(game.Players))]
	mTarget := messages[rand.Intn(len(messages))]

	if target.IsBanked || this.nextRoll {
		this.Bank()
		// Define the message to be spoken

		// Execute the say command with the message
		cmd := exec.Command("say", mTarget)

		// Run the command and check for errors
		go func() {
			err := cmd.Run()
			if err != nil {
				panic("Landon's Fault")
			}
		}()

		this.nextRoll = false
	}
}

func (this *NicolasPlayer) LastChance(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.GetNumNOTBankedPlayers() == 1 {
		this.nextRoll = true
	}
}

func (this *NicolasPlayer) RoundOver() {
	this.nextRoll = false
}

func (this *NicolasPlayer) GetName() string {
	return "He'll always try to stop me"
}
