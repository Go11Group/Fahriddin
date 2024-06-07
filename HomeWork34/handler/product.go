package handler

import (
	"encoding/json"
	"my_module/modul"
	"net/http"
	"strconv"
	"strings"
)

func (p *Products) ProductCreate(w http.ResponseWriter, r *http.Request) {
	product := modul.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		panic(err)
	}

	err = p.Product.Create(product)
	if err != nil {
		panic(err)
	}
}

func (p *Products) ProductRead(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/ReadUsers/"))
	if err != nil {
		panic(err)
	}
	user, err := p.Product.GetById(id)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		panic(err)
	}
}
func (p *Products) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	product := modul.Product{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		panic(err)
	}
	err = p.Product.Update(product, product.Id)
	if err != nil {
		panic(err)
	}
}

func (p *Products) ProductDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/DeleteUsers/"))
	if err != nil {
		panic(err)
	}

	err = p.Product.Delate(id)
	if err != nil {
		panic(err)
	}

	w.Write([]byte("Succes Delete storage"))
}
