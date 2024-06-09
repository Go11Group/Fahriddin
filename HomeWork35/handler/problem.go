package handler

import (
	"encoding/json"
	"my_module/modul"
	"net/http"
	"strings"
)

func (p *Problem) ConnectProblem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, "/Problems/")
		problem, err := p.Problem.GetById(id)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(problem)
		if err != nil {
			panic(err)
		}
	case "POST":
		problem := modul.Problem{}
		err := json.NewDecoder(r.Body).Decode(&problem)
		if err != nil {
			panic(err)
		}
		err = p.Problem.Create(problem)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Create storage"))
	case "PUT":
		problem := modul.Problem{}
		err := json.NewDecoder(r.Body).Decode(&problem)
		if err != nil {
			panic(err)
		}
		err = p.Problem.Update(problem)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Update storage"))
	case "DELETE":
		id := strings.TrimPrefix(r.URL.Path, "/Problems/")
		err := p.Problem.Delete(id)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Delete storage"))
	}
}
