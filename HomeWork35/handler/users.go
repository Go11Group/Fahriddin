package handler

import (
	"encoding/json"
	"my_module/modul"
	"net/http"
	"strings"
)

func (u *User) ConnectUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, "/Users/")
		user, err := u.User.GetById(id)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			panic(err)
		}
	case "POST":
		user := modul.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			panic(err)
		}
		err = u.User.Create(user)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Create storage"))
	case "PUT":
		user := modul.User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			panic(err)
		}
		err = u.User.Update(user)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Update storage"))
	case "DELETE":
		id := strings.TrimPrefix(r.URL.Path, "/Users/")
		err := u.User.Delete(id)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Succes Delete storage"))
	}
}
