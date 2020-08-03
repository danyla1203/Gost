package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	handlers      handlersMap
	middlewares   handlersMap
	staticDirName string
	staticFiles   []os.FileInfo
}

func (app App) Get(path string, handler handlerCallback) {
	app.handlers[path] = handler
}
func (app App) Use(path string, handler handlerCallback) {
	app.middlewares[path] = handler
}
func (app *App) Static(path string) {
	dirFiles, err := ioutil.ReadDir("./" + path)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("sdf")
	app.staticDirName = path[1:]
	app.staticFiles = dirFiles
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
	//handle static, if fist part of uri match static dir name
	//-> try to get file by second part of uri
	if splitedURI[0] == app.staticDirName && len(splitedURI) > 1 {
		fmt.Fprint(socket, getFile(app.staticDirName, splitedURI[1]))
		return
	}
	//get handler and matched pattern
	handler, pattern := GetHandler(app.handlers, splitedURI)
	if handler == nil {
		fmt.Fprint(socket, "Fuck, it's 404. Try another url, dude")
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
	//execute middlewares and handler
	for _, callback := range middlewares {
		callback(userRequest, userResponse)
	}
	handler(userRequest, userResponse)
}
