package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}).Methods("GET")
	return r
}



func Handler(w http.ResponseWriter, r *http.Request) {
	router := NewRouter()
	router.ServeHTTP(w, r)
}
