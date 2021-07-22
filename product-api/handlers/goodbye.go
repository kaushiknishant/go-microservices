package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	logger *log.Logger
}

func NewGoodBye(logger *log.Logger) *GoodBye {
	return &GoodBye{logger}
}

func (goodBye *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	goodBye.logger.Println("Good Bye World")
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Good Bye %s", response)
}
