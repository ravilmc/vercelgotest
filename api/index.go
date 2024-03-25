package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You have found light at the end of the tunnel!")
	})

	r.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Hello Your ID is %s", id)
	})

	r.HandleFunc("GET /api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := NewRouter()
	router.ServeHTTP(w, r)
}
