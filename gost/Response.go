package lib

import (
	"fmt"
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
	fmt.Fprint(res, file)
}
func (res Response) Redirect(location string) {
	res.Header().Set("Location", location)
	res.WriteHeader(http.StatusMovedPermanently)
}

func (res *Response) SetCookie(name, value string) {
	cookie := name + "=" + value + "; SameSite"
	res.Header().Add("Set-Cookie", cookie)
}
