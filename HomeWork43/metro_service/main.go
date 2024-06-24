package main

import (
	"my_module/api"
	"my_module/api/handler"
	"my_module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	cr := postgres.NewCardRepo(db)
	st := postgres.NewStationRepo(db)
	ter := postgres.NewTerminalRepo(db)
	tran := postgres.NewTransactionRepo(db)

	handler := handler.Handler{
		Card:        cr,
		Station:     st,
		Terminal:    ter,
		Transaction: tran,
	}

	router := api.Router(handler)

	router.Run(":8080")
}
