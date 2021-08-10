package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (product *Product) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r)
	return err.Decode(product)
}

// custom validation function for sku
func validateSKU(fl validator.FieldLevel) bool {
	//sku is of format abc-cdsf-asdfn
	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := regex.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
}

//function for validation
func (product *Product) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(product)
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (product *Products) ToJSON(w io.Writer) error {
	err := json.NewEncoder(w)
	return err.Encode(product)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// AddProduct add new product
func AddProduct(product *Product) {
	product.ID = getNextID()
	productList = append(productList, product)
}

// getNextID() will generate ID for the product added.
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// the UpdateProduct update the product in the product list
func UpdateProduct(id int, product *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	product.ID = id
	productList[pos] = product

	return nil
}

//custom error message
var ErrProductNotFound = fmt.Errorf("Product not found")

// Find the product in the Product List
func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
