package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/knagadevara/AkiraGames/GameOn/Blanks"
	Hangman "github.com/knagadevara/AkiraGames/GameOn/Hangman"
	staticfiles "github.com/knagadevara/AkiraGames/StaticFiles"
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
	enterSelection := string(utl.GetRune()(inpRdr))
	resp := staticfiles.BuildData()

	switch enterSelection {
	case "1":
		bl := Blanks.BlanksPlayer{}
		bl.Start(resp)
	case "2":
		cl := Hangman.CliffhangerPlayerData{}
		cl.Start()
	default:
		fmt.Printf("Please enter a valid number!!\n%v is not accepted\n", enterSelection)
	}

}
