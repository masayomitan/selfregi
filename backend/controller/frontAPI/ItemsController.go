package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"selfregi/middleware"
	"selfregi/model"
	"strconv"

	"github.com/gorilla/mux"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	middleware.CorsMiddleware(w, r)
	vars := mux.Vars(r)
	itemId, _ := strconv.Atoi(vars["item_id"])
	i := &model.Item{}
	res, _ := i.GetItemDisplayOn(context.Background(), itemId)
	json.NewEncoder(w).Encode(res)
}

func GetItemsByCategoryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId, _ := strconv.Atoi(vars["category_id"])
	middleware.CorsMiddleware(w, r)
	i := &model.Items{}
	res, _ := i.GetItemsDisplayOnFromCategoryId(context.Background(), categoryId)
	json.NewEncoder(w).Encode(res)
}
