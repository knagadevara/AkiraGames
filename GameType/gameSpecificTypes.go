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

type GameName interface {
	GetGameName()
}

type BlanksPlayerData struct {
	Puzzel      *Country // Holds the data of Puzzel
	CrypticWord string   // Holds the display word
	GuessWord   string   // Current Guess Word
	Name        string   // Game Name
	IsCorrect   bool     // If the guess is correct will be set to true
	TryCount    int      // Total number of tries
}

type HangmanPlayerData struct {
	BPD                  *BlanksPlayerData
	CurrentGuessedLetter rune           //  collects the current guessed letter
	LettersInWord        map[rune][]int // makes a map of runes with its index
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
