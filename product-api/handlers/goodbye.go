package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GoodBye is a simple handler
type GoodBye struct {
	logger *log.Logger
}

//NewGoodBye creates a new GoodBye handler with the given logger
func NewGoodBye(logger *log.Logger) *GoodBye {
	return &GoodBye{logger}
}

//SeverHTTP implements the go http.Handler interface
//https://golang.org/pkg/net/http/#Handler
func (goodBye *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	goodBye.logger.Println("Handle Good Bye Requests")

	//read the body
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		goodBye.logger.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}
	//write the response
	fmt.Fprintf(rw, "Good Bye %s", response)
}
