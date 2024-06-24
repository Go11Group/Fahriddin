package main

import (
	"my_module/api"
	"my_module/api/handler"
	"my_module/storage/postgres"
)


func main(){
	db,err := postgres.ConnectPostgres()

	if err != nil{
		panic(err)
	}

	us := postgres.NewUsersRepo(db)

	handler := handler.Handler{
		User: us,
	}

	router := api.Router(handler)

	router.Run(":8080")
}