package api

import (
	"fmt"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	router.HandleFunc("POST /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World post")
	})

	router.HandleFunc("GET /api/{version}/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Version: %s", r.PathValue("version"))
	})

	return router
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := NewRouter()
	router.ServeHTTP(w, r)
}
