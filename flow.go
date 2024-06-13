package AkiraGames

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/knagadevara/AkiraGames/utl"
)

func (c *Country) SetCountry() *Country {
	apiBaseUrl := "https://countriesnow.space/api/"
	apiVersion := "v0.1"
	apiResource := "/countries/capital"
	resource_string := apiBaseUrl + apiVersion + apiResource
	CountryResp := utl.LoadGameData[CountryApiResp]("GET", resource_string, "/Users/snagadev/go/src/AkiraGames/StaticFiles/GameJSON/Countries.json")
	if CountryResp.Error != "" {
		log.Fatalln("Unable to Get Data")
	}
	return utl.GetRandItem(CountryResp.Data)
}
func (c *Country) GetCountry() *string { return &c.Name }
func (c *Country) GetCapital() *string { return &c.Capital }
func (c *Country) GetISO() *string     { return &c.ISO2 }

func displayGames() {
	fmt.Println("Please Select a Game")
	fmt.Println("1. Hangman")
}

func PopGame() {
	displayGames()
	inpRdr := bufio.NewReader(os.Stdin)
	enterSelection := string(utl.GetRune()(inpRdr))

	switch enterSelection {
	case "1":
		cl := CliffhangerPlayerData{}
		cl.Start()
	default:
		fmt.Printf("Please enter a valid number!!\n%v is not accepted\n", enterSelection)
	}

}
