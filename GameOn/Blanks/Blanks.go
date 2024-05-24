package Blanks

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	GameType "github.com/knagadevara/AkiraGames/GameType"
	"github.com/knagadevara/AkiraGames/utl"
)

type BlanksPlayer GameType.BlanksPlayerData

type HangMan interface {
	DisplayGameState() *BlanksPlayer
	GetInput() *BlanksPlayer
	GetGussWord(Countries []GameType.Country) *BlanksPlayer // Selects a random country and its capital.
	Match() *BlanksPlayer                                   // Matches the gussed word with puzzled word.
	MakePuzzleWord() *BlanksPlayer
}

// When called takes input and gives a rune.
func (h *BlanksPlayer) GetInput() *BlanksPlayer {
	fmt.Printf("Please Input your Guess! :\t")
	inpRdr := bufio.NewReader(os.Stdin)
	h.GuessWord = utl.GetString()(inpRdr)
	return h
}

func (h *BlanksPlayer) GetGussWord(Countries []GameType.Country) *BlanksPlayer {
	h.Puzzel = &Countries[rand.Intn(len(Countries))]
	h.Puzzel.Name = strings.ToLower(h.Puzzel.Name)
	return h
}

func (h *BlanksPlayer) MakePuzzleWord() *BlanksPlayer {
	wc := len(h.Puzzel.Name) - 1
	blanks := wc / 2
	crossedString := []rune(h.Puzzel.Name)
	for i := 0; i <= blanks; i++ {
		crossedString[rand.Intn(wc)] = '_'
	}
	h.CrypticWord = string(crossedString)
	return h
}

func (h *BlanksPlayer) Match() *BlanksPlayer {
	if h.Puzzel.Name == h.GuessWord {
		h.IsCorrect = true
		fmt.Println("SUPER!!!")
		fmt.Printf("Country:\t%v\nCapital:\t%v\nA.K.A:\t\t%v\n", h.Puzzel.Name, h.Puzzel.Capital, h.Puzzel.ISO2)
	} else {
		h.TryCount += 1
		h.IsCorrect = false
	}
	return h
}

func (h *BlanksPlayer) DisplayGameState() *BlanksPlayer {
	insigNia := "\t\t=====| * |=====\t\t"
	header := insigNia + " H A N G M A N " + insigNia
	footer := insigNia + " * + - | - + * " + insigNia
	pedastal := "===\n=====\n======="
	pole := "\n||\n||\n||\n||"
	hanger := "\t============"
	hanggedMan := "|\n|\nO\n/M\\\nA\nH\n>.<"
	fmt.Printf("%v", header)
	fmt.Printf("Guess Me!!!! %v\n", h.CrypticWord)
	switch h.TryCount {
	case 1:
		fmt.Printf("%v\n", pedastal)
	case 2:
		fmt.Println("HINT!!!!:\t\t", h.Puzzel.ISO2)
		fmt.Printf("%v\n", pole)
		fmt.Printf("\t\t%v\n", pedastal)
	case 3:
		fmt.Println("HINT!!!!:\t\t", h.Puzzel.Capital)
		fmt.Printf("%v", hanger)
		fmt.Printf("%v\n", pole)
		fmt.Printf("\t\t%v\n", pedastal)
	case 4:
		fmt.Printf("%v%v", hanger, hanggedMan)
		fmt.Printf("%v\n", pole)
		fmt.Printf("\t\t%v\n", pedastal)
	default:
		fmt.Println()
	}
	fmt.Printf("%v", footer)
	return h
}

func (h *BlanksPlayer) GamePlay() {
	for !(h.IsCorrect) {
		if h.DisplayGameState().
			GetInput().
			Match().TryCount > 4 {
			break
		}
	}
}

func (h *BlanksPlayer) Start(resp GameType.CountryApiResp) {
	h.GetGussWord(resp.Rastra).
		MakePuzzleWord().
		GamePlay()
}
