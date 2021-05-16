package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
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
