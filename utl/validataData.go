package utl

import (
	"log"
	"os"
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
