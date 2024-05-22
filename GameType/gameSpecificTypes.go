package GameType

/*
###########
GenericGame
###########
*/

type TextGame interface {
	DisplayGameState()
	GameOn()
}

/*
#######
HANGMAN
#######
*/
type HangmanPlayerData struct {
	Puzzel                 *Country
	GuessWord, CrypticWord string
	IsCorrect              bool
	TryCount               int8
}

type Country struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
	ISO2    string `json:"iso2"`
	ISO3    string `json:"iso3"`
}

type HangmanApiResp struct {
	Error  string    `json:"error"`
	Msg    string    `json:"msg"`
	Rastra []Country `json:"data"`
}
