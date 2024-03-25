package api

import (
	"fmt"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World post")
	})

	router.HandleFunc("GET /{version}/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Version: %s", r.PathValue("version"))
	})

	return router
}

func Handler(w http.ResponseWriter, r *http.Request) {
	api := http.NewServeMux()

	api.Handle("/api/", http.StripPrefix("/api", NewRouter()))

	api.ServeHTTP(w, r)
}
