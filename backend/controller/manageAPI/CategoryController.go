package controller

import (
	"context"
	"encoding/json"
	// "fmt"
	"net/http"
	"selfregi/ent"
	"selfregi/model"
	"selfregi/middleware"
)
type Categories struct {
  entCategories[] *ent.Categories
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	middleware.CorsMiddleware(w)
	a := &model.Categories{}
	res, _ := a.GetCategories(context.Background())
	json.NewEncoder(w).Encode(res)
}