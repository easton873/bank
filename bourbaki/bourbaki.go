package bourbaki

import (
	"bank/bank"
	"encoding/base64"
	"math/rand"
	"os/exec"
)

type NicolasPlayer struct {
	bank.PlayerControls
	nextRoll bool
}

// Decodes a Base64 string into the original message
func decodeMessage(encoded string) string {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic("Failed to decode message: " + err.Error())
	}
	return string(data)
}

func (this *NicolasPlayer) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	// Base64-encoded messages
	encodedMessages := []string{
		"YWxsIG91ciBiYXNlIGlzIGJlbG9uZyB0byB1cw==",
		"SSdtIHNvIHdlYWsw",
		"bmljbw==",
		"ZWFzdCBpcyB1cA==",
		"SSBhbSBjbGFuY3k=",
		"c3Vkb3cgd29vZG8=",
		"RWFzdG9uLCBJJ20gbm90IGZlZWxpbmcgd2VsbA==",
	}

	// Decode the messages before using them
	messages := make([]string, len(encodedMessages))
	for i, encoded := range encodedMessages {
		messages[i] = decodeMessage(encoded)
	}

	// Select a target player and a random message
	target := game.Players[rand.Intn(len(game.Players))]
	mTarget := messages[rand.Intn(len(messages))]

	// If the target is banked or we need the next roll, proceed
	if target.IsBanked || this.nextRoll {
		this.Bank()

		// Execute the say command with the message
		cmd := exec.Command("say", mTarget)

		// Run the command asynchronously
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
