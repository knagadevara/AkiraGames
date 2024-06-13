package AkiraGames

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/knagadevara/AkiraGames/utl"
)

type Blanks interface {
	DisplayGameState() *BlanksPlayer
	GetInput() *BlanksPlayer
	GetGussWord(Countries []Country) *BlanksPlayer // Selects a random country and its capital.
	Match() *BlanksPlayer                          // Matches the gussed word with puzzled word.
	MakePuzzleWord() *BlanksPlayer
}

// When called takes input and gives a rune.
func (h *BlanksPlayer) GetInput() *BlanksPlayer {
	fmt.Printf("Please Input your Guess! :\t")
	inpRdr := bufio.NewReader(os.Stdin)
	h.GuessWord = utl.GetString()(inpRdr)
	return h
}

func (h *BlanksPlayer) MakePuzzleWord(Countries []Country) *BlanksPlayer {
	h.Puzzel.SetCountry()
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
	header := insigNia + "\t" + h.Name + "\t" + insigNia
	footer := insigNia + " * + - | - + * " + insigNia
	fmt.Printf("Guess Me??? >>>> %v", h.CrypticWord)
	fmt.Printf("%v", header)
	switch h.TryCount {
	case 2:
	case 3:
		fmt.Println("HINT!!!!:\t\t", h.Puzzel.ISO2)
	case 4:
		fmt.Println("HINT!!!!:\t\t", h.Puzzel.Capital)
	case 5:
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

func (h *BlanksPlayer) Start(resp CountryApiResp) {
	h.MakePuzzleWord(resp.Data).
		GamePlay()
}
