package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
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

func (h *HangManPlayer) GetGussWord(Countries []GameType.Country) *HangManPlayer {
	h.Puzzel = &Countries[rand.Intn(len(Countries))]
	h.Puzzel.Name = strings.ToLower(h.Puzzel.Name)
	return h
}

func (h *HangManPlayer) MakeDashes() *HangManPlayer {
	h.CrypticWord = strings.Repeat("-", len(h.Puzzel.Name))
	return h
}

// Makes a pay which adds the index of similar numbers
func (h *HangManPlayer) CountOfLetters() *HangManPlayer {
	runeCounter := make(map[rune][]int)
	for ix, v := range h.CrypticWord {
		runeCounter[v] = append(runeCounter[v], ix)
	}
	h.LettersInWord = runeCounter
	return h
}

func (h *HangManPlayer) CheckAndRevealWord() *HangManPlayer {
	crossword := []rune(h.CrypticWord)
	ix, ok := h.LettersInWord[h.CurrentGuessedLetter]
	if ok {
		for _, v := range ix {
			crossword[v] = h.CurrentGuessedLetter
		}
	} else {
		log.Panicln("Wrong Guess!!!")
		h.TryCount += 1
	}
	h.CrypticWord = string(crossword)
	if strings.ContainsRune(h.CrypticWord, '-') {
		h.IsCorrect = false
	} else {
		h.IsCorrect = true
	}
	return h
}

func (h *HangManPlayer) DisplayGameState() *HangManPlayer {
	insigNia := "\t\t=====| * |=====\t\t"
	header := insigNia + " H A N G M A N " + insigNia
	footer := insigNia + " * + - | - + * " + insigNia
	log.Println(header)
	log.Printf("Guess Me??? >>>> %v", h.CrypticWord)
	log.Println(footer)
	pedastal := "===\n=====\n======="
	pole := "\n||\n||\n||\n||"
	hanger := "\t============"
	hanggedMan := "|\n|\nO\n/M\\\nA\nH\n>.<"
	fmt.Printf("%v", header)
	fmt.Printf("Guess Me!!!! %v\n", h.CrypticWord)
	switch h.TryCount {
	case 2:
		fmt.Printf("%v\n", pedastal)
	case 3:
		fmt.Println("HINT!!!!:\t\t", h.Puzzel.ISO2)
		fmt.Printf("%v\n", pole)
		fmt.Printf("\t\t%v\n", pedastal)
	case 4:
		fmt.Println("HINT!!!!:\t\t", h.Puzzel.Capital)
		fmt.Printf("%v", hanger)
		fmt.Printf("%v\n", pole)
		fmt.Printf("\t\t%v\n", pedastal)
	case 5:
		fmt.Printf("%v%v", hanger, hanggedMan)
		fmt.Printf("%v\n", pole)
		fmt.Printf("\t\t%v\n", pedastal)
	default:
		fmt.Println()
	}
	fmt.Printf("%v", footer)
	return h
}

func (h *HangManPlayer) GamePlay() {
	for !(h.IsCorrect) {
		if h.DisplayGameState().
			GetInput().
			CheckAndRevealWord().
			TryCount > len(h.CrypticWord) {
			break
		}
	}
}

func (h *HangManPlayer) Start(resp GameType.CountryApiResp) {
	h.GetGussWord(resp.Rastra).
		MakeDashes().
		CountOfLetters().
		GamePlay()
}
