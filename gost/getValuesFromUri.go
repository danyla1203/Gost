package lib

import "strconv"

func getValuesFromUri(uri []string, path []string) map[string]int {
	uriValues := make(map[string]int)
	// if uri == "/" and path == "/"
	if len(uri) == 1 && len(uri[0]) == 0 && len(path) == 1 && len(path[0]) == 0 {
		return uriValues
	}
	for index, _ := range uri {
		if path[index][0:1] == ":" {
			valueName := path[index][1:]
			uriValues[valueName], _ = strconv.Atoi(uri[index])
		}
	}
	return uriValues
}
