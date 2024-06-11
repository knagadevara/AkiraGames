package GameType

/*
###########
GenericGame
###########
*/

type TextGameIface interface {
	DisplayGameState()
	GameOn()
}

type Country struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
	ISO2    string `json:"iso2"`
	ISO3    string `json:"iso3"`
}

type CountryApiResp struct {
	Error string    `json:"error"`
	Msg   string    `json:"msg"`
	Data  []Country `json:"data"`
}
