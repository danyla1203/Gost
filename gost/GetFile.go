package lib

import (
	"io/ioutil"
	"log"
)

func GetFile(dirName, fileName string) string {
	file, err := ioutil.ReadFile(dirName + "/" + fileName)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(file[:])
}
