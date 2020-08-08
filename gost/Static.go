package lib

import (
	"io/ioutil"
	"log"
	"os"
)

type Static struct {
	staticDirName string
	staticFiles   []os.FileInfo
}

func (static *Static) ServeFile(path string) {
	dirFiles, err := ioutil.ReadDir("./" + path)
	if err != nil {
		log.Print(err)
		return
	}
	static.staticDirName = path[1:]
	static.staticFiles = dirFiles
}
