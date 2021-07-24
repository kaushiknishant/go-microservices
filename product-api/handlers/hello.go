package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	logger *log.Logger
}

//NewHello creates a new hello handler with the given logger
func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

//SeverHTTP implements the go http.Handler interface
//https://golang.org/pkg/net/http/#Handler
func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hello.logger.Println("Handle Hello Requests")

	//read the body
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		hello.logger.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}
	//write the response
	fmt.Fprintf(rw, "Hello %s", response)
}
