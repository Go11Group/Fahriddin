package handler

import (
	"encoding/json"
	"my_module/modul"
	"net/http"
	"strconv"
	"strings"
)

func (u *Users) UserCreate(w http.ResponseWriter, r *http.Request) {
	user := modul.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	err = u.User.Create(user)
	if err != nil {
		panic(err)
	}
}

func (u *Users) UserRead(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/ReadUsers/"))
	if err != nil {
		panic(err)
	}
	user, err := u.User.GetById(id)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		panic(err)
	}
}
func (u *Users) UserUpdate(w http.ResponseWriter, r *http.Request) {
	user := modul.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	err = u.User.Update(user, user.Id)
	if err != nil {
		panic(err)
	}
}

func (u *Users) UserDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/DeleteUsers/"))
	if err != nil {
		panic(err)
	}

	err = u.User.Delate(id)

	if err != nil {
		panic(err)
	}

	w.Write([]byte("Succes Delete storage"))
}
