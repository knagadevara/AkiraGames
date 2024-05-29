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

/*
#######
HANGMAN
#######
*/

type Name string                  // Game Name
type Puzzel string                // Holds the data of Puzzel
type CrypticWord string           // Holds the display word
type GuessWord string             // Current Guess Word
type IsCorrect bool               // If the guess is correct will be set to true
type TryCount int                 // Total number of tries
type CurrentGuessedLetter rune    //  collects the current guessed letter
type LettersInWord map[rune][]int // makes a map of runes with its index
type PreviousLetters map[rune]bool

type PuzzleIface interface {
	GetPuzzle()
	SetPuzzle()
}

type GuessWordIface interface {
	GetGuessWord()
	SetGuessWord()
}
type CrypticWordIface interface {
	GetCrypticWord()
	SetCrypticWord()
}
type CurrentGuessedLetterIface interface {
	GetCurrentGuessedLetter()
	SetCurrentGuessedLetter()
}
type CurrentTryCountIface interface {
	GetCurrentTryCount()
	SetCurrentTryCount()
}

type GameNameIface interface {
	GetGameName()
	SetGameName()
}

type IsCorrectIface interface {
	GetIsCorrect()
	SetIsCorrect()
}

type PreviousWordsIface interface {
	GetPreviousWords()
	SetPreviousWords()
}

type PreviousLettersIface interface {
	GetPreviousLetters()
	SetPreviousLetters()
}

type LettersInWordIface interface {
	GetLettersInWord()
	SetLettersInWord()
}
type BlanksPlayerData struct {
	Puzzel        *Country // Holds the data of Puzzel
	CrypticWord   string   // Holds the display word
	GuessWord     string   // Current Guess Word
	Name          string   // Game Name
	IsCorrect     bool     // If the guess is correct will be set to true
	TryCount      int      // Total number of tries
	PreviousWords map[string]bool
}

type HangmanPlayerData struct {
	BPD                  *BlanksPlayerData
	CurrentGuessedLetter rune           //  collects the current guessed letter
	LettersInWord        map[rune][]int // makes a map of runes with its index
	PreviousLetters      map[rune]bool
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
