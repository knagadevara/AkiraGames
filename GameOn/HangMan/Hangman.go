package Hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	GameType "github.com/knagadevara/AkiraGames/GameType"
	"github.com/knagadevara/AkiraGames/utl"
)

type HangManPlayer GameType.HangmanPlayerData

type HangMan interface {
	DisplayGameState() *HangManPlayer
	GetInput() *HangManPlayer
	GetGussWord(Countries []GameType.Country) *HangManPlayer // Selects a random country and its capital.
	MatchReveal() *HangManPlayer                             // Matches the gussed word with puzzled word.
	MakePuzzleWord() *HangManPlayer
}

// Done
// When called takes input and gives a rune.
func (h *HangManPlayer) GetInput() *HangManPlayer {
	fmt.Printf("Please Input your Guess! :\t")
	inpRdr := bufio.NewReader(os.Stdin)
	h.CurrentGuessedLetter = utl.GetRune()(inpRdr)
	return h
}

// Done
// Show space as space, replace everything else with '-'
func (h *HangManPlayer) MakePuzzel(Countries []GameType.Country) *HangManPlayer {
	h.BPD.Puzzel = utl.GetCountry(Countries)
	tmpRunes := make([]rune, len(h.BPD.Puzzel.Name))
	for i, v := range h.BPD.Puzzel.Name {
		if v == ' ' {
			tmpRunes[i] = ' '
		} else {
			tmpRunes[i] = '_'
		}
	}
	h.BPD.CrypticWord = string(tmpRunes)
	return h
}

// Done
// Makes a pay which adds the index of similar numbers
func (h *HangManPlayer) CountOfLetters() *HangManPlayer {
	runeCounter := make(map[rune][]int)
	for ix, v := range h.BPD.Puzzel.Name {
		runeCounter[v] = append(runeCounter[v], ix)
	}
	h.LettersInWord = runeCounter
	fmt.Println(h.LettersInWord)
	return h
}

// Check If the gussed letter exists in LettersInWord map
// else increases try-count and adds the gussed letter to PreviousLetters.
// Checks if all letters are completed
func (h *HangManPlayer) CheckAndRevealWord() *HangManPlayer {
	crossword := []rune(h.BPD.CrypticWord)
	ix, ok := h.LettersInWord[h.CurrentGuessedLetter]
	if ok {
		for _, v := range ix {
			crossword[v] = h.CurrentGuessedLetter
		}
	} else {
		h.PreviousLetters[h.CurrentGuessedLetter] = true
		h.BPD.TryCount += 1
		fmt.Println("Wrong Guess!!!")
	}
	h.BPD.CrypticWord = string(crossword)
	if strings.ContainsRune(h.BPD.CrypticWord, '-') {
		h.BPD.IsCorrect = false
	} else {
		h.BPD.IsCorrect = true
	}
	return h
}

func (h *HangManPlayer) DisplayGameState() *HangManPlayer {

	insigNia := "\t\t=====| * |=====\t\t"
	header := insigNia + "\t" + h.BPD.Name + "\t" + insigNia
	footer := insigNia + " * + - | - + * " + insigNia
	fmt.Printf("Guess Me??? >>>> %v", h.BPD.CrypticWord)
	fmt.Printf("%v", header)
	switch h.BPD.TryCount {
	case 2:
	case 3:
		fmt.Println("HINT!!!!:\t\t", h.BPD.Puzzel.ISO2)
	case 4:
		fmt.Println("HINT!!!!:\t\t", h.BPD.Puzzel.Capital)
	case 5:
	default:
		fmt.Println()
	}
	fmt.Printf("%v", footer)
	return h
}

func (h *HangManPlayer) GamePlay() {
	for !(h.BPD.IsCorrect) {
		if h.DisplayGameState().
			GetInput().
			CheckAndRevealWord().
			BPD.TryCount > len(h.BPD.CrypticWord) {
			break
		}
	}
}

func (h *HangManPlayer) Start(resp GameType.CountryApiResp) {
	h.MakePuzzel(resp.Rastra).
		CountOfLetters().
		GamePlay()
}
