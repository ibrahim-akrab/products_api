package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ibrahim-akrab/products_api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(rw, r)
		return
	case http.MethodPost:
		p.AddProduct(rw, r)
		return
	case http.MethodPut:
		p.UpdateProduct(rw, r)
		return
	default:
		rw.WriteHeader(http.StatusNotImplemented)
	}

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetProducts()
	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle product addition")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
	d, _ := json.Marshal(prod)
	rw.Write(d)

}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle product update")
}
