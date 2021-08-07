package handlers

import (
	"context"
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

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	//add the data
	data.AddProduct(&prod)
}

func (product Products) UpdateProduct(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	product.logger.Println("Handle PUT Product",id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func(product Products) MiddlewareProductValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r* http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
	
		if err != nil {
			product.logger.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}
				// add the product to the context
				ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
				r = r.WithContext(ctx)
		
				// Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(rw, r)
	})
}
