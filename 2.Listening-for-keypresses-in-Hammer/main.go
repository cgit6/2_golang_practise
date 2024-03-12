package main

import (
	"fmt"
	"myapp/game"
)

func main() {
	playAgain := true

	for playAgain {
		// game.Play()
		playAgain = game.GetYesOrNo("你想再玩一次嗎 (y/n)?")
	}

	fmt.Println("")
	fmt.Println("Goodbye.")
}
