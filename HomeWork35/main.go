package main

import (
	"my_module/handler"
	"my_module/storage/postgres"
	"net/http"
)

func main() {
	db, err := postgres.ConnectDB()

	if err != nil {
		panic(err)
	}

	user := postgres.NewUserRepo(db)
	problem := postgres.NewProblemRepo(db)
	solved_problem := postgres.NewSolved_problemsRepo(db)

	mux := handler.NewHandler(handler.User{User: user}, handler.Problem{Problem: problem},handler.Solved_problem{Solved_problem: solved_problem})
	http.ListenAndServe(":8080", mux)
}
