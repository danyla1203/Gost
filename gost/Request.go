package lib

import (
	"log"
	"net/http"
)

type Request struct {
	*http.Request
	UrlParts map[string]int
	Body     map[string]string
}

func (req *Request) SetParams() {
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
