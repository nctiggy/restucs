package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"encoding/json"
	"github.com/nctiggy/restucs/chasis"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/chasis", ChasisIndex)
	router.HandleFunc("/chasis/{chasisId}", ChasisShow)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {

	chasis := ChasisList{
		Chasis{Name: "Chasis1"},
		Chasis{Name: "Chasis2"},
	}

	json.NewEncoder(w).Encode(ChasisList)

}

func ChasisIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Chasis Index")

}

func ChasisShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	chasisId := vars["chasisId"]
	fmt.Fprintln(w, "Chasis show:", chasisId)

}
