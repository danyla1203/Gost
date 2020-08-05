package lib

import (
	"errors"
	"strings"
)

func GetHandler(handlers []Handler, uri []string, method string) (Handler, error) {
	for _, handlerObj := range handlers {
		if handlerObj.method == method {
			splitedPattern := strings.Split(handlerObj.path, "/")[1:]
			//check path is matching pattern
			isSuitable := CheckPath(uri, splitedPattern)
			if isSuitable {
				return handlerObj, nil
			}
		}
	}
	return Handler{}, errors.New("No handler")
}
