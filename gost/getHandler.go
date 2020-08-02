package lib

import (
	"strings"
)

func GetHandler(handlers handlersMap, uri []string) (handlerCallback, []string) {
	for handlerPattern, handlerFunc := range handlers {
		splitedPattern := strings.Split(handlerPattern, "/")[1:]
		//check path is matching pattern
		isSuitable := CheckPath(uri, splitedPattern)
		if isSuitable {
			return handlerFunc, splitedPattern
		}
	}
	return nil, nil
}
