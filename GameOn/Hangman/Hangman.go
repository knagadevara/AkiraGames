package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knagadevara/AkiraGames/utl"
)

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

func (g GameName) SetGameName(name string) *GameName {
	g = GameName(name)
	return &g
}

func (Is IsCorrect) SetIsCorrect(tf bool) IsCorrect { return IsCorrect(tf) }

func (c CrypticWord) SetCrypticWord(word GuessWord) *CrypticWord {
	tmpRunes := make([]rune, len(word))
	for i, v := range word {
		if v == ' ' {
			tmpRunes[i] = ' '
		} else {
			tmpRunes[i] = '_'
		}
	}
	c = CrypticWord(tmpRunes)
	return &c
}

func (g GuessWord) SetGuessWord(c *Country) *GuessWord {
	ptrStng := c.GetCountry()
	g = GuessWord(*ptrStng)
	return &g
}

// Checks if all letters are completed
func (Ic IsCorrect) CheckIfCorrect(cw CrypticWord) IsCorrect {
	if strings.ContainsRune(string(cw), '-') {
		return Ic.SetIsCorrect(true)
	} else {
		return Ic.SetIsCorrect(false)
	}
}

// Makes an array of indexs of all the letters in word in Map
func (l LettersInWord) SetLettersInWord(g GuessWord) *LettersInWord {
	if l == nil {
		l = make(LettersInWord)
	}
	for ix, v := range g {
		l[Letter(v)] = append(l[Letter(v)], ix)
	}
	return &l
}

// Check If the gussed letter exists in LettersInWord map
func (crypt CrypticWord) CheckIfLetterExists(l LettersInWord, guessLetter Letter) *CrypticWord {
	crossword := []Letter(crypt)
	ix, ok := l[guessLetter]
	if ok {
		for _, v := range ix {
			crossword[v] = guessLetter
		}
		crypt = CrypticWord(crossword)
		return &crypt
	}
	return &crypt
}

// Adds the word to Previously gussed list.
func (pl PreviousLetters) SetPreviousLetters(guessLetter Letter) *PreviousLetters {
	if pl == nil {
		pl = make(PreviousLetters)
	}
	_, ok := pl[guessLetter]
	if ok {
		fmt.Println("Already Guessed Letter!!")
	} else {
		pl[guessLetter] = true
	}
	return &pl
}

// else increases try-count and adds the gussed letter to PreviousLetters.
func (bl TryCount) SetTryCount(i int) *TryCount {
	bl += TryCount(i)
	return &bl
}

// When called takes input and gives a rune.
func (l Letter) SetLetter() *Letter {
	fmt.Printf("Please Input your Guess! :\t")
	inpRdr := bufio.NewReader(os.Stdin)
	l = Letter(utl.GetRune()(inpRdr))
	return &l
}

// Show Hangman Status "StaticFiles/hangmanStates/1"
func (Cf CliffhangerPlayerData) PrintHangman(filePath string) {
	utl.CheckFileExists(filePath)
	flBuff := utl.LoadFile(filePath)
	fmt.Println(string(flBuff))
}

// Display Header and Footer onlyh once
func (Cf CliffhangerPlayerData) PrintHeader(Header string) { fmt.Printf("%v\n\n", Header) }
func (Cf CliffhangerPlayerData) PrintFooter(Footer string) { fmt.Printf("%v\n\n", Footer) }

// Score Board.
func (Cf CliffhangerPlayerData) DisplayGameState() *CliffhangerPlayerData {
	fmt.Printf("Guess Me>>>>\t\t%v\n", string(*Cf.CrypticWord))
	fmt.Printf("Tries\t\t\t\t%v\n", *Cf.TryCount)
	return &Cf
}

func (Cf *CliffhangerPlayerData) Initiate() *CliffhangerPlayerData {
	var g GameName
	Cf.Name = g.SetGameName("Cliffhanger")
	c := &Country{}
	Cf.Country = c.SetCountry()
	var gw GuessWord
	Cf.GuessWord = gw.SetGuessWord(Cf.Country)
	lw := make(LettersInWord, len(*Cf.GuessWord))
	Cf.LettersInWord = lw.SetLettersInWord(*Cf.GuessWord)
	var Ic IsCorrect
	Cf.IsCorrect = Ic.SetIsCorrect(false)
	Cf.LastGusessCorrect = IsCorrect(false)
	var Tc TryCount
	Cf.TryCount = Tc.SetTryCount(0)
	var cw CrypticWord
	Cf.CrypticWord = cw.SetCrypticWord(*Cf.GuessWord)
	var l Letter
	Cf.CurrentGuessedLetter = &l
	var pvl PreviousLetters
	Cf.PreviousLetters = &pvl
	Cf.PrintHeader(Cf.InsigNia(string(*Cf.Name)))
	return Cf
}

func (Cf *CliffhangerPlayerData) GamePlay() *CliffhangerPlayerData {
	wordLen := TryCount(len(*Cf.CrypticWord))
	for !(Cf.IsCorrect) {
		Cf = Cf.DisplayGameState()
		Cf.CurrentGuessedLetter = Cf.CurrentGuessedLetter.SetLetter()
		Cf.CrypticWord = Cf.CrypticWord.CheckIfLetterExists(*Cf.LettersInWord, *Cf.CurrentGuessedLetter)
		Cf.PreviousLetters = Cf.PreviousLetters.SetPreviousLetters(*Cf.CurrentGuessedLetter)
		if strings.Contains(string(*Cf.CrypticWord), string(*Cf.CurrentGuessedLetter)) {
			Cf.TryCount = Cf.TryCount.SetTryCount(0)
		} else {
			if *Cf.TryCount <= 11 {
				flNm := fmt.Sprintf("StaticFiles/hangmanStates/%v", *Cf.TryCount)
				Cf.PrintHangman(flNm)
			}
			Cf.TryCount = Cf.TryCount.SetTryCount(1)
		}
		Cf.IsCorrect = Cf.IsCorrect.CheckIfCorrect(*Cf.CrypticWord)
		if *Cf.TryCount <= 12 {
			if *Cf.TryCount > wordLen+4 {
				fmt.Printf("Maximum Tries Reached!!!\n")
				break
			}
		}
	}
	return Cf
}

func (Cf *CliffhangerPlayerData) InsigNia(mark string) string {
	insigNia := "\t\t=====| * |=====\t\t"
	return fmt.Sprintf(insigNia + "\t" + mark + "\t" + insigNia)
}

func (Cf *CliffhangerPlayerData) Start() {
	footer := " * + - | - + * "
	Cf.Initiate().
		GamePlay().
		PrintFooter(Cf.InsigNia(footer))
}
