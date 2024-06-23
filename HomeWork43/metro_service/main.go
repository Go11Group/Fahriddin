package main

import (
	"my_module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	
	cr := postgres.NewCardRepo(db)
	st := postgres.NewStationRepo(db)
	
}
