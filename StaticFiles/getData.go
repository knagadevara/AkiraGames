package staticfiles

import (
	"fmt"

	"github.com/knagadevara/AkiraGames/GameType"
	"github.com/knagadevara/AkiraGames/utl"
)

func BuildData() GameType.CountryApiResp {
	apiBaseUrl := "https://countriesnow.space/api/"
	apiVersion := "v0.1"
	apiResource := "/countries/capital"
	resource_string := fmt.Sprintf(apiBaseUrl + apiVersion + apiResource)
	return utl.LoadGameData[GameType.CountryApiResp]("GET", resource_string, "StaticFiles/GameJSON/Countries.json")
}
