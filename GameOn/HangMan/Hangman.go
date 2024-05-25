package Hangman

import (
	"bufio"
	"fmt"
	"log"
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

// When called takes input and gives a rune.
func (h *HangManPlayer) GetInput() *HangManPlayer {
	fmt.Printf("Please Input your Guess! :\t")
	inpRdr := bufio.NewReader(os.Stdin)
	h.CurrentGuessedLetter = utl.GetRune()(inpRdr)
	return h
}

func (h *HangManPlayer) MakePuzzel(Countries []GameType.Country) *HangManPlayer {
	h.BPD.Puzzel = utl.GetCountry(Countries)
	h.BPD.CrypticWord = strings.Repeat("-", len(h.BPD.Puzzel.Name))
	return h
}

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

func (h *HangManPlayer) CheckAndRevealWord() *HangManPlayer {
	crossword := []rune(h.BPD.CrypticWord)
	ix, ok := h.LettersInWord[h.CurrentGuessedLetter]
	if ok {
		for _, v := range ix {
			crossword[v] = h.CurrentGuessedLetter
		}
	} else {
		log.Println("Wrong Guess!!!")
		h.BPD.TryCount += 1
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
	header := insigNia + " H A N G M A N " + insigNia
	footer := insigNia + " * + - | - + * " + insigNia
	log.Println(header)
	log.Printf("Guess Me??? >>>> %v", h.BPD.CrypticWord)
	log.Println(footer)
	pedastal := "===\n=====\n======="
	pole := "\n||\n||\n||\n||"
	hanger := "============"
	hanggedMan := "|\n|\nO\n/M\\\nA\nH\n>.<"
	fmt.Printf("%v", header)
	fmt.Printf("Guess Me!!!! %v\n", h.BPD.CrypticWord)
	switch h.BPD.TryCount {
	case 2:
		fmt.Printf("%v\n", pedastal)
	case 3:
		fmt.Println("HINT!!!!:\t\t", h.BPD.Puzzel.ISO2)
		fmt.Printf("%v\n", pole)
		fmt.Printf("%v\n", pedastal)
	case 4:
		fmt.Println("HINT!!!!:\t\t", h.BPD.Puzzel.Capital)
		fmt.Printf("%v\n", hanger)
		fmt.Printf("%v\n", pole)
		fmt.Printf("%v\n", pedastal)
	case 5:
		fmt.Printf("%v\t\t%v\n", hanger, hanggedMan)
		fmt.Printf("%v\n", pole)
		fmt.Printf("%v\n", pedastal)
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
