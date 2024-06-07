package handler

import (
	"my_module/postgres"
	"net/http"
)

type Users struct {
	User *postgres.UserRepo
}

type Products struct {
	Product *postgres.ProductRepo
}

func NewHandler(users Users, products Products) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /CreateUsers/", users.UserCreate)
	mux.HandleFunc("GET /ReadUsers/",users.UserRead)
	mux.HandleFunc("PUT /UpdateUsers/",users.UserUpdate)
	mux.HandleFunc("DELETE /DeleteUsers/",users.UserDelete)

	mux.HandleFunc("POST /CreateProducts/", products.ProductCreate)
	mux.HandleFunc("GET /ReadProducts/", products.ProductRead)
	mux.HandleFunc("PUT /UpdateProducts/", products.ProductUpdate)
	mux.HandleFunc("DELETE /DeleteProducts/", products.ProductDelete)

	return mux
}
