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

	mux.HandleFunc("/CreateUsers/", users.UserCreate)
	mux.HandleFunc("/ReadUsers/",users.UserRead)
	mux.HandleFunc("/UpdateUsers/",users.UserUpdate)
	mux.HandleFunc("/DeleteUsers/",users.UserDelete)

	mux.HandleFunc("/CreateProducts/", products.ProductCreate)
	mux.HandleFunc("/ReadProducts/", products.ProductRead)
	mux.HandleFunc("/UpdateProducts/", products.ProductUpdate)
	mux.HandleFunc("/DeleteProducts/", products.ProductDelete)

	return mux
}
