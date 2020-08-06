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
	Body     map[string]string
}
type Response struct {
	http.ResponseWriter
}

func (res Response) SendFile(path string) {
	//by default, files to send must be placed in assets dir
	//TODO: Get correct mime type
	file, mimeType := GetFile("/assets/" + path)
	res.Header().Set("Content-Type", mimeType)
	fmt.Fprint(res, file)
}

type handlerCallback func(r *Request, res *Response)

type App struct {
	handlers      []Handler
	middlewares   []Middleware
	staticDirName string
	staticFiles   []os.FileInfo
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
func (app *App) Static(path string) {
	dirFiles, err := ioutil.ReadDir("./" + path)
	if err != nil {
		log.Print(err)
		return
	}
	app.staticDirName = path[1:]
	app.staticFiles = dirFiles
}

func MakeApp() App {
	//init empty maps
	app := App{
		handlers:    []Handler{},
		middlewares: []Middleware{},
	}
	return app
}

func setParams(req *Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal("something wrong with read data from body")
	}
	reducedMap := map[string]string{}
	for key, value := range req.PostForm {
		reducedMap[key] = value[0]
	}
	req.Body = reducedMap
}

func (app App) ServeHTTP(socket http.ResponseWriter, request *http.Request) {
	splitedURI := strings.Split(request.RequestURI, "/")[1:]
	//handle static, if first part of request uri match static dir name
	//-> try to get file by second part of request uri
	if splitedURI[0] == app.staticDirName && len(splitedURI) > 1 {
		//TODO: Get correct mime type
		file, mimeType := GetFile("/" + app.staticDirName + "/" + splitedURI[1])
		socket.Header().Set("Content-Type", mimeType)
		fmt.Fprint(socket, file)
		return
	}
	//get handler and matched pattern
	handler, err := GetHandler(app.handlers, splitedURI, request.Method)
	if err != nil {
		fmt.Fprint(socket, "Fuck, it's 404. Try another url, dude")
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
	setParams(userRequest)
	//execute middlewares and handler
	for _, callback := range middlewares {
		callback(userRequest, userResponse)
	}
	handler.handler(userRequest, userResponse)
}
