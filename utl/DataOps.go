package utl

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/knagadevara/AkiraGames/GameType"
)

// Load data from file.
func LoadGameData[T any](httpMethod, apiURL, fileName string) T {
	var file *os.File

	fileInfo := CheckFileExists(fileName)
	defer file.Close()

	if fileInfo != nil && fileInfo.Size() > 0 {

		file = OperateFile(fileName, os.O_RDONLY, 0655)
		data := DecodeFileToStruct[T](file)
		return data

	} else {

		log.Printf("Creating %s...\n", fileName)
		resp := CallApi(httpMethod, apiURL)
		defer resp.Body.Close()

		file = OperateFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0655)
		_ = WriteToFile(file, resp.Body)

		data := DecodeFileToStruct[T](file)
		return data
	}
}

// When called takes input and gives a String.
func GetString() func(inpRdr *bufio.Reader) string {
	return func(inpRdr *bufio.Reader) string {
		word, err := inpRdr.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		return strings.ToLower(strings.TrimSpace(word))
	}
}

// When called takes input and gives a Rune.
func GetRune() func(inpRdr *bufio.Reader) rune {
	return func(inpRdr *bufio.Reader) rune {
		r, _, err := inpRdr.ReadRune()
		if err != nil {
			log.Fatalln(err)
		}
		return r
	}
}

// Selects a random word
func GetCountry(Countries []GameType.Country) *GameType.Country {
	Puzzel := &Countries[rand.Intn(len(Countries))]
	Puzzel.Name = strings.ToLower(Puzzel.Name)
	Puzzel.Capital = strings.ToLower(Puzzel.Capital)
	return Puzzel
}

// func DisplayGameState[T any]() T {
// 	insigNia := "\t\t=====| * |=====\t\t"
// 	header := insigNia + " H A N G M A N " + insigNia
// 	footer := insigNia + " * + - | - + * " + insigNia
// 	log.Println(header)
// 	log.Printf("Guess Me??? >>>> %v", h.CrypticWord)
// 	log.Println(footer)
// 	pedastal := "===\n=====\n======="
// 	pole := "\n||\n||\n||\n||"
// 	hanger := "============"
// 	hanggedMan := "|\n|\nO\n/M\\\nA\nH\n>.<"
// 	fmt.Printf("%v", header)
// 	fmt.Printf("Guess Me!!!! %v\n", h.CrypticWord)
// 	switch h.TryCount {
// 	case 2:
// 		fmt.Printf("%v\n", pedastal)
// 	case 3:
// 		fmt.Println("HINT!!!!:\t\t", h.Puzzel.ISO2)
// 		fmt.Printf("%v\n", pole)
// 		fmt.Printf("%v\n", pedastal)
// 	case 4:
// 		fmt.Println("HINT!!!!:\t\t", h.Puzzel.Capital)
// 		fmt.Printf("%v\n", hanger)
// 		fmt.Printf("%v\n", pole)
// 		fmt.Printf("%v\n", pedastal)
// 	case 5:
// 		fmt.Printf("%v\t\t%v\n", hanger, hanggedMan)
// 		fmt.Printf("%v\n", pole)
// 		fmt.Printf("%v\n", pedastal)
// 	default:
// 		fmt.Println()
// 	}
// 	fmt.Printf("%v", footer)
// 	return h
// }
