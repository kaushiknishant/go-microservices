package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hello.logger.Println("Hello World")
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", response)
}
