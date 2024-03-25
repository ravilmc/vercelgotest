package api

import (
	"net/http"

	"github.com/ravilmc/vercelgotest/business/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := handler.NewHandler()
	handler.ServeHTTP(w, r)
}
