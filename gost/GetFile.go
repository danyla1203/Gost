package lib

import (
	"io/ioutil"
	"log"
	"os"
)

func GetFile(fileName string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	file, err := ioutil.ReadFile(path + fileName)

	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(file[:])
}
