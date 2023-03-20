package controller

import (
		"net/http"
		"context"
		"encoding/json"
		"selfregi/ent"
		"selfregi/model"
		"selfregi/middleware"
)

type Categories struct {
    entCategories[] *ent.Categories
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
		middleware.CorsMiddleware(w, r)
		a := &model.Categories{}
		res, _ := a.GetCategoriesDisplayOn(context.Background())
		json.NewEncoder(w).Encode(res)
}