package handlers

import (
	"encoding/json"
	"log"
	"microservies/product-api/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := json.NewEncoder(rw).Encode(lp)
	if err != nil {
		http.Error(rw, "Error marshalling products.", http.StatusInternalServerError)
	}
}
