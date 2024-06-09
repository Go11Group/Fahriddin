package handler

import (
	"my_module/storage/postgres"

	"github.com/gorilla/mux"
)

type Solved_problem struct{
	Solved_problem *postgres.Solved_problemsRepo
}

type User struct {
	User *postgres.UserRepo
}

type Problem struct {
	Problem *postgres.ProblemRepo
}

func NewHandler(user User, problem Problem, sp Solved_problem) *mux.Router {
	m := mux.NewRouter()

	m.HandleFunc("/Users/", user.ConnectUser)
	m.HandleFunc("/Problems/",problem.ConnectProblem)

	return m
}
