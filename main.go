package main

import (
	"net/http"

	"github.com/ravilmc/vercelgotest/api"
)

func main() {

	server := http.Server{
		Addr:    ":3000",
		Handler: &Handler{},
	}

	server.ListenAndServe()
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.Handler(w, r)
}
