package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/knagadevara/AkiraGames/GameOn/Blanks"
	cliffhanger "github.com/knagadevara/AkiraGames/GameOn/CliffHanger"
	Hangman "github.com/knagadevara/AkiraGames/GameOn/HangMan"
	staticfiles "github.com/knagadevara/AkiraGames/StaticFiles"
	"github.com/knagadevara/AkiraGames/utl"
)

func displayGames() {
	fmt.Println("Please Select a Game")
	fmt.Println("1. Hangman")
	fmt.Println("2. Blanks")
	fmt.Println("3. CliffHanger")

}

func main() {
	displayGames()
	inpRdr := bufio.NewReader(os.Stdin)
	enterSelection := string(utl.GetRune()(inpRdr))
	resp := staticfiles.BuildData()

	switch enterSelection {
	case "1":
		hm := Hangman.HangManPlayer{}
		hm.Start(resp)
	case "2":
		bl := Blanks.BlanksPlayer{}
		bl.Start(resp)
	case "3":
		cl := cliffhanger.CliffhangerPlayerData{}
		cl.Start()
	default:
		fmt.Printf("Please enter a valid number!!\n%v is not accepted\n", enterSelection)
	}

}
