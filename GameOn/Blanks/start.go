package blanks

import (
	"fmt"

	"github.com/knagadevara/AkiraGames/GameType"
	"github.com/knagadevara/AkiraGames/utl"
)

func start() {
	apiBaseUrl := "https://countriesnow.space/api/"
	apiVersion := "v0.1"
	apiResource := "/countries/capital"
	resource_string := fmt.Sprintf(apiBaseUrl + apiVersion + apiResource)
	resp := utl.LoadGameData[GameType.HangmanApiResp]("GET", resource_string, "../StaticFiles/GameJSON/Countries.json")
	hangman := HangManPlayer{}
	hangman.
		GetGussWord(resp.Rastra).
		MakePuzzleWord().
		GameOn()
}
