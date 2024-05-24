package main

import (
	"bufio"
	"fmt"
	"os"

	bl "github.com/knagadevara/AkiraGames/GameOn/Blanks"
	hg "github.com/knagadevara/AkiraGames/GameOn/HangMan"
	"github.com/knagadevara/AkiraGames/utl"
)

func displayGames() {
	fmt.Println("Please Select a Game")
	fmt.Println("1. Hangman")
	fmt.Println("2. Blanks")
}

func main() {
	displayGames()
	inpRdr := bufio.NewReader(os.Stdin)
	enterSelection := utl.GetRune()(inpRdr)
	switch enterSelection {
	case 1:
		hg.Start()
	case 2:
		bl.Start()
	}

}
