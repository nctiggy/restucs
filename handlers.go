package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {

fmt.Fprintln(w, "Welcome!")

}

func ChasisIndex(w http.ResponseWriter, r *http.Request) {

	chasislist := Chasislist{
		Chasis{Name: "Chasis1"},
		Chasis{Name: "Chasis2"},
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(chasislist); err != nil {
		panic(err)
	}

}

func ChasisShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	chasisId := vars["chasisId"]
	fmt.Fprintln(w, "Chasis show:", chasisId)

}
