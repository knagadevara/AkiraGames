package main

import (
	"fmt"

	GameType "github.com/knagadevara/AkiraGames/GameType"
	"github.com/knagadevara/AkiraGames/HangMan/GameOn"
	utl "github.com/knagadevara/AkiraGames/utl"
)

var resp GameType.HangmanApiResp

func init() {
	apiBaseUrl := "https://countriesnow.space/api/"
	apiVersion := "v0.1"
	apiResource := "/countries/capital"
	resource_string := fmt.Sprintf(apiBaseUrl + apiVersion + apiResource)
	resp = utl.LoadGameData[GameType.HangmanApiResp]("GET", resource_string, "../StaticFiles/HangMan/Countries.json")
}

func main() {
	hangman := GameOn.HangManPlayer{}
	hangman.
		GetGussWord(resp.Rastra).
		MakePuzzleWord().
		GameOn()
}
