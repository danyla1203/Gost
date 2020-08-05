package lib

import (
	"strconv"
	"strings"
)

func GetValuesFromUri(uri []string, path string) map[string]int {
	uriValues := make(map[string]int)
	splitedPath := strings.Split(path, "/")[1:]
	// if uri == "/" and path == "/"
	if len(uri) == 1 && len(uri[0]) == 0 && len(splitedPath) == 1 && len(splitedPath[0]) == 0 {
		return uriValues
	}
	for index, _ := range uri {
		if splitedPath[index][0:1] == ":" {
			valueName := splitedPath[index][1:]
			uriValues[valueName], _ = strconv.Atoi(uri[index])
		}
	}
	return uriValues
}
