package controller

import (
		"fmt"
		"log"
		"net/http"
		"encoding/json"
		// "github.com/gorilla/mux"
		"selfregi/middleware"
		"selfregi/component"
		"selfregi/entity"
		"selfregi/usecase"
		"io/ioutil"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	fmt.Fprintf(w, "get pructsss")
}

type Item struct {
		*entity.ItemEntity
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Print(len(body))
	if len(body) > 0 {
		fmt.Println("no null")

		middleware.CorsMiddleware(w, r)
		ctx := r.Context()
		var item = entity.ItemEntity{}
		err := component.DecodeJSONBody(w, r, &item)
		if err != nil {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		fmt.Println(item)
		usecase.ItemCreate(ctx, &item)
		// json.NewEncoder(w).Encode(item)
		fmt.Println("ok")
}
}