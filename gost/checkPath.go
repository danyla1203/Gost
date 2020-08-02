package lib

import (
	"regexp"
)

func checkPath(uri []string, path []string) bool {
	// if uri == "/" and path == "/"
	if len(uri) == 1 && len(uri[0]) == 0 && len(path) == 1 && len(path[0]) == 0 {
		return true
	}
	if len(uri) != len(path) {
		return false
	}
	for index, _ := range uri {
		if path[index][0:1] == ":" {
			continue
		}

		isMatchByRegExp, _ := regexp.MatchString(path[index], uri[index])
		if !isMatchByRegExp {
			return false
		} else {
			continue
		}
	}
	return true
}