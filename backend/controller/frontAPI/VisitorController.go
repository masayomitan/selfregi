package controller

import (
	// "fmt"
	"net/http"
	"selfregi/usecase"
	// "selfregi/model"
	// "github.com/gorilla/mux"
)


func VisitorStore (w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	usecase.Create(ctx)
}