package AkiraGames

type CountryApiResp struct {
	Error string    `json:"error"`
	Msg   string    `json:"msg"`
	Data  []Country `json:"data"`
}

type Country struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
	ISO2    string `json:"iso2"`
	ISO3    string `json:"iso3"`
}
type CountryIface interface {
	SetCountry() *Country
	GetCountry() *string
	GetCapital() *string
	GetISO() *string
}

type GameName string // Game Name
type GameNameIface interface {
	SetGameName(name string) *GameName
}

type CrypticWord string // Holds the display word
type CrypticWordIface interface {
	SetCrypticWord(word GuessWord) *CrypticWord
	CheckIfLetterExists(l LettersInWord, guessLetter Letter) *CrypticWord
}

type GuessWord string // Current Guess Word
type GuessWordIface interface {
	SetGuessWord(c *Country) *GuessWord
}

type IsCorrect bool // If the guess is correct will be set to true
type IsCorrectIface interface {
	SetIsCorrect(tf bool) IsCorrect
	CheckIfCorrect(cw CrypticWord) IsCorrect
}

type LettersInWord map[Letter][]int // makes a map of runes with its index
type LettersInWordIface interface {
	SetLettersInWord(s string) *LettersInWord
}

type Letter rune //  collects the current guessed letter
type LetterIface interface {
	SetLetter() *Letter
}

type TryCount int // Total number of tries
type TryCountIface interface {
	SetTryCount(i int) *TryCount
}
type PreviousLetters map[Letter]bool
type PreviousLettersIface interface {
	SetPreviousLetters(guessLetter Letter) *PreviousLetters
}

type Cliffhanger interface {
	GameNameIface
	CountryIface
	CrypticWordIface
	GuessWordIface
	IsCorrectIface
	PreviousLettersIface
	TryCountIface
	LetterIface
	LettersInWordIface
	Start() *CliffhangerPlayerData
	GamePlay() *CliffhangerPlayerData
}

type CliffhangerPlayerData struct {
	Name                 *GameName
	Country              *Country
	CrypticWord          *CrypticWord
	GuessWord            *GuessWord
	IsCorrect            IsCorrect
	TryCount             *TryCount
	CurrentGuessedLetter *Letter
	LettersInWord        *LettersInWord
	PreviousLetters      *PreviousLetters
	LastGusessCorrect    IsCorrect
}

type BlanksPlayer struct {
	Puzzel        *Country // Holds the data of Puzzel
	CrypticWord   string   // Holds the display word
	GuessWord     string   // Current Guess Word
	Name          string   // Game Name
	IsCorrect     bool     // If the guess is correct will be set to true
	TryCount      int      // Total number of tries
	PreviousWords map[string]bool
}
