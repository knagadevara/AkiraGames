package main

import (
	"bufio"
	"fmt"
	"os"

	Hangman "github.com/knagadevara/AkiraGames/GameOn/Hangman"
	"github.com/knagadevara/AkiraGames/utl"
)

func displayGames() {
	fmt.Println("Please Select a Game")
	fmt.Println("1. Hangman")
}

func main() {
	displayGames()
	inpRdr := bufio.NewReader(os.Stdin)
	enterSelection := string(utl.GetRune()(inpRdr))

	switch enterSelection {
	case "1":
		cl := Hangman.CliffhangerPlayerData{}
		cl.Start()
	default:
		fmt.Printf("Please enter a valid number!!\n%v is not accepted\n", enterSelection)
	}

}
