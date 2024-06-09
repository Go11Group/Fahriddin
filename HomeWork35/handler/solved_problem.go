package handler

import (
	"encoding/json"
	"my_module/modul"
	"net/http"
	"strings"
)

func (s *Solved_problem)ConnectSolved_problem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, "/ReadUsers/")
		sp, err := s.Solved_problem.GetById(id)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(sp)
		if err != nil {
			panic(err)
		}
	case "POST":
		sp := modul.Solved_problem{}
		err := json.NewDecoder(r.Body).Decode(&sp)
		if err != nil {
			panic(err)
		}
		err = s.Solved_problem.Create(sp)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Create storage"))
	case "PUT":
		sp := modul.Solved_problem{}
		err := json.NewDecoder(r.Body).Decode(&sp)
		if err != nil {
			panic(err)
		}
		err = s.Solved_problem.Update(sp)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Update storage"))
	case "DELETE":
		id := strings.TrimPrefix(r.URL.Path, "/DeleteUsers/")
		err := s.Solved_problem.Delete(id)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Delete storage"))
	}
}
