package data

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
	"time"
)

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

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", ValidateSku)
	return validate.Struct(p)

}
func ValidateSku(fl validator.FieldLevel) bool {
	regx := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := regx.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProductById(id)
	if err != nil {
		return err
	}
	p.ID = id
	products[pos] = p
	return nil
}

var ErrProductNotFound = errors.New("Product not found")

func findProductById(id int) (*Product, int, error) {
	for i, v := range products {
		if v.ID == id {
			return v, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func AddProduct(p *Product) {
	p.ID = getNextID()

	products = append(products, p)

}

func getNextID() int {
	ln := products[len(products)-1]
	return ln.ID + 1
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return products
}

var products = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.99,
		SKU:         "abc3221",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       2.00,
		SKU:         "asd123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
