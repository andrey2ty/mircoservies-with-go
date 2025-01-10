package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"microservies/product-api/data"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Could not convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT product", id)
	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	p.l.Printf("Added product %#v", prod)
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

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post product")
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	p.l.Printf("Added product %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := json.NewEncoder(rw).Encode(lp)
	if err != nil {

		http.Error(rw, "Error marshalling products.", http.StatusInternalServerError)
	}
}
