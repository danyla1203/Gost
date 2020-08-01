package lib

import (
	"fmt"
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

type App struct {
	handlers map[string]func(request Request, response Response)
}

func (app App) Get(path string, handler func(request Request, response Response)) {
	app.handlers[path] = handler
}

func MakeApp() App {
	app := App{handlers: map[string]func(request Request, response Response){}}
	return app
}

func (app App) ServeHTTP(socket http.ResponseWriter, request *http.Request) {
	splitedURI := strings.Split(request.RequestURI, "/")[1:]
	for key, value := range app.handlers {
		splitedPath := strings.Split(key, "/")[1:]
		//check path is matching pattern
		isSuitable := checkPath(splitedURI, splitedPath)
		if isSuitable {
			valuesFromUri := getValuesFromUri(splitedURI, splitedPath)
			//create own request, response struct
			userRequest := Request{request, valuesFromUri}
			userResponse := Response{socket}
			value(userRequest, userResponse)
			return
		}
	}
	fmt.Fprint(socket, "Fuck, nothing here. 404 motherfucker")
}
