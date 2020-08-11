package lib

import (
	"io/ioutil"
	"log"
	"os"
)

func GetFile(fileName string) (string, string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	file, err := ioutil.ReadFile(path + fileName)

	if err != nil {
		log.Fatal(err)
		return "", ""
	}
	//TODO: how to get this fucking mime?
	return string(file[:]), ""
}
