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

type Puzzle interface {
	GetPuzzle()
}
type GuessWord interface {
	GetGuessWord()
}
type CrypticWord interface {
	GetCrypticWord()
}
type CurrentGuessedLetter interface {
	GetCurrentGuessedLetter()
}
type CurrentTryCount interface {
	GetCurrentTryCount()
}

type BlanksPlayerData struct {
	Puzzel                 *Country
	GuessWord, CrypticWord string
	IsCorrect              bool
	TryCount               int
}

type HangmanPlayerData struct {
	Puzzel               *Country
	CrypticWord          string
	IsCorrect            bool
	TryCount             int
	CurrentGuessedLetter rune
	LettersInWord        map[rune][]int
}

type Country struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
	ISO2    string `json:"iso2"`
	ISO3    string `json:"iso3"`
}

type CountryApiResp struct {
	Error  string    `json:"error"`
	Msg    string    `json:"msg"`
	Rastra []Country `json:"data"`
}
