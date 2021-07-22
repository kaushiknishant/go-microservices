package main

import (
	"net/http"
	//"log"
	"fmt"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// log.Println("Hello world")
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Bad Request", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", response)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Bad Request", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Good Bye %s", res)
	})

	http.ListenAndServe(":9090", nil)
}
