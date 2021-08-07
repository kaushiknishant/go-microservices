package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
func (product *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
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

func (product Products) UpdateProduct(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err1 := strconv.Atoi(vars["id"])
	if err1 != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	product.logger.Println("Handle PUT Product",id)
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
