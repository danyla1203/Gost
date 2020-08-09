package lib

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type handlerCallback func(r *Request, res *Response)

type App struct {
	Static
	handlers    []Handler
	middlewares []Middleware
}

type Middleware struct {
	path     string
	callback handlerCallback
}
type Handler struct {
	method  string
	path    string
	handler handlerCallback
}

func (app *App) Get(path string, handler handlerCallback) {
	handlerObj := Handler{
		method:  "GET",
		path:    path,
		handler: handler,
	}
	app.handlers = append(app.handlers, handlerObj)
}
func (app *App) Post(path string, handler handlerCallback) {
	handlerObj := Handler{
		method:  "POST",
		path:    path,
		handler: handler,
	}
	app.handlers = append(app.handlers, handlerObj)
}
func (app *App) Use(path string, handler handlerCallback) {
	middleware := Middleware{
		path:     path,
		callback: handler,
	}
	app.middlewares = append(app.middlewares, middleware)
}

func MakeApp() App {
	//init empty maps
	app := App{
		handlers:    []Handler{},
		middlewares: []Middleware{},
	}
	return app
}

func (app App) ServeHTTP(socket http.ResponseWriter, request *http.Request) {
	splitedURI := strings.Split(request.RequestURI, "/")[1:]
	//handle static, if first part of request uri match static dir name
	//-> try to get file by second part of request uri
	if splitedURI[0] == app.staticDirName && len(splitedURI) > 1 {
		//TODO: Get correct mime type
		file, mimeType := GetFile("/" + app.staticDirName + "/" + splitedURI[1])
		socket.Header().Set("Content-Type", mimeType)
		socket.Header().Set("Cache-Control", "max-age=216000; must-revalidate")
		fmt.Fprint(socket, file)
		return
	}
	//get handler and matched pattern
	handler, err := GetHandler(app.handlers, splitedURI, request.Method)
	if err != nil {
		log.Fatal(fmt.Fprint(socket, "Fuck, it's 404. Try another url, dude"))
		return
	}
	//get vars from uri by pattern
	valuesFromUri := GetValuesFromUri(splitedURI, handler.path)
	//get all middlewares for current uri
	middlewares := GetMiddlewares(app.middlewares, splitedURI)
	//create modified req, resp objects
	userRequest := &Request{
		Request:  request,
		UrlParts: valuesFromUri,
	}
	userResponse := &Response{socket}
	userRequest.SetParams()
	//execute middlewares and handler
	for _, callback := range middlewares {
		callback(userRequest, userResponse)
	}
	handler.handler(userRequest, userResponse)
}
