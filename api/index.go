package api

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := NewHandler()
	handler.ServeHTTP(w, r)
}
