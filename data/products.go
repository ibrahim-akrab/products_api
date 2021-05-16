package data

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {

	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

func GetProducts() Products {
	return ProductsList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	ProductsList = append(ProductsList, p)
}

func getNextId() int64 {
	lp := ProductsList[len(ProductsList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int64, p *Product) error {
	return nil
}

var ProductsList = []*Product{
	{
		ID:          1,
		Name:        "Mocha",
		Description: "A pretty good mocha",
		Price:       3.22,
		SKU:         "abc",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Latte",
		Description: "An awesome latte",
		Price:       2.99,
		SKU:         "def",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
