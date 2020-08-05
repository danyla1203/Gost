package lib

import (
	"strings"
)

func GetMiddlewares(middlewares []Middleware, uri []string) []handlerCallback {
	selectedMiddlewares := []handlerCallback{}
	for _, middlewareObj := range middlewares {
		splitedMiddlewarePattern := strings.Split(middlewareObj.path, "/")[1:]
		if IsMiddleware(uri, splitedMiddlewarePattern) {
			selectedMiddlewares = append(selectedMiddlewares, middlewareObj.callback)
		}
	}
	return selectedMiddlewares
}
