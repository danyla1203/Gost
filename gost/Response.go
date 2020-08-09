package lib

import (
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func (res Response) SendFile(path string) {
	//by default, files to send must be placed in assets dir
	//TODO: Get correct mime type
	file, mimeType := GetFile("/assets/" + path)
	res.Header().Set("Content-Type", mimeType)
	log.Fatal(fmt.Fprint(res, file))
}

func (res *Response) SetCookie(name, value string) {
	cookie := name + "=" + value + "; SameSite"
	res.Header().Add("Set-Cookie", cookie)
}
