package lib

import (
	"io/ioutil"
	"log"
	"net/http"
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
	return string(file[:]), http.DetectContentType(file)
}
