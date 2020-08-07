package lib

import (
	"reflect"
	"regexp"
)

func CheckPath(uri []string, path []string) bool {
	// if uri == "/" and path == "/"
	if len(path) == 1 && len(path[0]) == 0 {
		return reflect.DeepEqual(path, uri)
	}
	if len(uri) != len(path) {
		return false
	}
	for index, _ := range uri {
		if path[index][0:1] == ":" || path[index] == "*" {
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
