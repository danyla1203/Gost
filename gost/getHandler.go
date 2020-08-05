package lib

import (
	"errors"
	"strings"
)

func GetHandler(handlers []Handler, uri []string) (Handler, error) {
	for _, handlerObj := range handlers {
		splitedPattern := strings.Split(handlerObj.path, "/")[1:]
		//check path is matching pattern
		isSuitable := CheckPath(uri, splitedPattern)
		if isSuitable {
			return handlerObj, nil
		}
	}
	return Handler{}, errors.New("No handler")
}
