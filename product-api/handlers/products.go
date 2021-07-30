package handlers

import (
	"log"
	"net/http"

	"github.com/kaushiknishant/go-microservices/product-api/data"
)

// Products is a http.Handler
type Products struct {
	logger *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (product *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		product.GetProducts(rw, r)
		return
	}
	// handle addition of data into List
	if r.Method == http.MethodPost {
		product.addProduct(rw, r)
		return
	}

	//catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (product *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	product.logger.Println("Handle GET Products")

	// fetch the products from the datastore
	productList := data.GetProducts()

	// serialize the list to JSON
	err := productList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to give result ", http.StatusInternalServerError)
	}
}

// addProduct added new product in data store
func (product *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	product.logger.Println("Handle POST Product")

	// Product defines the structure for an API product
	prod := &data.Product{}
	// serialize the JSON to list
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	product.logger.Printf("Prod: %#v", prod)

	//add the data
	data.AddProduct(prod)
}
