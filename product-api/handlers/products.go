package handlers

import (
	"log"
	"net/http"

	"github.com/kaushiknishant/go-microservices/product-api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (product *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		product.GetProducts(rw, r)
		return
	}
	// handle an update 
	

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (product *Products) GetProducts(rw http.ResponseWriter,r *http.Request){
	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	
	if err != nil {
		http.Error(rw, "unable to give result ", http.StatusInternalServerError)
	}
}
