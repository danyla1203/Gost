package lib

import (
	"net/http"
	"strings"
)

type Request struct {
	*http.Request
	UrlParts map[string]int
}
type Response struct {
	http.ResponseWriter
}
type handlerCallback func(r *Request, res *Response)
type handlersMap map[string]func(request *Request, response *Response)

type App struct {
	handlers    handlersMap
	middlewares handlersMap
}

func (app App) Get(path string, handler handlerCallback) {
	app.handlers[path] = handler
}
func (app App) Use(path string, handler handlerCallback) {
	app.middlewares[path] = handler
}

func MakeApp() App {
	//init empty maps
	app := App{
		handlers:    handlersMap{},
		middlewares: handlersMap{},
	}
	return app
}

func (app App) ServeHTTP(socket http.ResponseWriter, request *http.Request) {
	splitedURI := strings.Split(request.RequestURI, "/")[1:]
	//get handler and matched pattern
	handler, pattern := GetHandler(app.handlers, splitedURI)
	if handler == nil {
		return
	}
	//get vars from uri by pattern
	valuesFromUri := GetValuesFromUri(splitedURI, pattern)
	//get all middlewares for current uri
	middlewares := GetMiddlewares(app.middlewares, splitedURI)
	//create modified req, resp objects
	userRequest := &Request{
		Request:  request,
		UrlParts: valuesFromUri,
	}
	userResponse := &Response{socket}
	//execute middlewares
	for _, callback := range middlewares {
		callback(userRequest, userResponse)
	}
	handler(userRequest, userResponse)
}
