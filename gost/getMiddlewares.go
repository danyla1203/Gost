package lib

import (
	"strings"
)

func GetMiddlewares(middlewares handlersMap, uri []string) []handlerCallback {
	selectedMiddlewares := []handlerCallback{}
	for middlewarePattern, callback := range middlewares {
		splitedMiddlewarePattern := strings.Split(middlewarePattern, "/")[1:]
		if IsMiddleware(uri, splitedMiddlewarePattern) {
			selectedMiddlewares = append(selectedMiddlewares, callback)
		}
	}
	return selectedMiddlewares
}
