package cliffhanger

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knagadevara/AkiraGames/utl"
)

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
	SetIsCorrect(tf bool) *IsCorrect
	CheckIfCorrect(cw CrypticWord) *IsCorrect
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
	IsCorrect            *IsCorrect
	TryCount             *TryCount
	CurrentGuessedLetter *Letter
	LettersInWord        *LettersInWord
	PreviousLetters      *PreviousLetters
}

func (c *Country) SetCountry() *Country {
	apiBaseUrl := "https://countriesnow.space/api/"
	apiVersion := "v0.1"
	apiResource := "/countries/capital"
	resource_string := apiBaseUrl + apiVersion + apiResource
	CountryResp := utl.LoadGameData[struct {
		Error     string
		Msg       string
		Countries []Country
	}]("GET", resource_string, "StaticFiles/GameJSON/Countries.json")
	if CountryResp.Error != "" {
		log.Fatalln("Unable to Get Data")
	}
	return utl.GetRandItem(CountryResp.Countries)
}
func (c *Country) GetCountry() *string { return &c.Name }
func (c *Country) GetCapital() *string { return &c.Capital }
func (c *Country) GetISO() *string     { return &c.ISO2 }

func (g GameName) SetGameName(name string) *GameName {
	g = GameName(name)
	return &g
}

func (Is IsCorrect) SetIsCorrect(tf bool) *IsCorrect {
	Is = IsCorrect(tf)
	return &Is
}

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
func (Ic IsCorrect) CheckIfCorrect(cw CrypticWord) *IsCorrect {
	if strings.ContainsRune(string(cw), '-') {
		Ic = *Ic.SetIsCorrect(false)
		return &Ic
	} else {
		Ic = IsCorrect(true)
		return &Ic
	}
}

// Makes an array of indexs of all the letters in word in Map
func (l LettersInWord) SetLettersInWord(g GuessWord) *LettersInWord {
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
	fmt.Println("Wrong Guess!!!")
	return &crypt
}

// Adds the word to Previously gussed list.
func (pl PreviousLetters) SetPreviousLetters(guessLetter Letter) *PreviousLetters {
	pl[guessLetter] = true
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

func (Cf *CliffhangerPlayerData) DisplayGameState() *CliffhangerPlayerData {
	insigNia := "\t\t=====| * |=====\t\t"
	header := insigNia + "\t" + string(*Cf.Name) + "\t" + insigNia
	footer := insigNia + " * + - | - + * " + insigNia
	fmt.Printf("%v", header)
	fmt.Printf("Guess Me??? >>>> %v", Cf.CrypticWord)
	fmt.Printf("%v", footer)
	return Cf
}

func (Cf *CliffhangerPlayerData) Initiate() *CliffhangerPlayerData {
	Cf.Name = Cf.Name.SetGameName("Cliffhanger")
	Cf.Country = Cf.Country.SetCountry()
	Cf.GuessWord = Cf.GuessWord.SetGuessWord(Cf.Country)
	Cf.LettersInWord = Cf.LettersInWord.SetLettersInWord(*Cf.GuessWord)
	Cf.IsCorrect = Cf.IsCorrect.SetIsCorrect(false)
	Cf.TryCount = Cf.TryCount.SetTryCount(0)
	Cf.CrypticWord = Cf.CrypticWord.SetCrypticWord(*Cf.GuessWord)
	return Cf
}

func (Cf *CliffhangerPlayerData) GamePlay() {
	word := TryCount(len(*Cf.CrypticWord))
	for !(*Cf.IsCorrect) {
		Cf.CurrentGuessedLetter = Cf.DisplayGameState().CurrentGuessedLetter.SetLetter()
		Cf.CrypticWord = Cf.CrypticWord.CheckIfLetterExists(*Cf.LettersInWord, *Cf.CurrentGuessedLetter)
		Cf.PreviousLetters = Cf.PreviousLetters.SetPreviousLetters(*Cf.CurrentGuessedLetter)
		Cf.IsCorrect = Cf.IsCorrect.CheckIfCorrect(*Cf.CrypticWord)
		Cf.TryCount = Cf.TryCount.SetTryCount(1)
		if *Cf.TryCount > word {
			break
		}
	}
}

func (Cf *CliffhangerPlayerData) Start() {
	Cf.Initiate().
		GamePlay()
}
