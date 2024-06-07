package main

import (
	"my_module/handler"
	"my_module/postgres"
	"net/http"
)

func main(){
	db,err := postgres.ConnectDB()
	if err != nil{
		panic(err)
	}
	products := postgres.NewProductRepo(db)
	users := postgres.NewUsersRepo(db)

	mux := handler.NewHandler(handler.Users{User: users},handler.Products{Product: products})

	
	http.ListenAndServe(":8080",mux)
}