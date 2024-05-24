package utl

import (
	"bufio"
	"log"
	"os"
	"strings"
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
