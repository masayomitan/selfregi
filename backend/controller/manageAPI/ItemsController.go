package controller

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	// "github.com/gorilla/mux"
	"selfregi/middleware"
	"selfregi/component"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	fmt.Fprintf(w, "get pructsss")
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
}
var books []Book

func PostItem(w http.ResponseWriter, r *http.Request) {
	middleware.CorsMiddleware(w)
	var book Book
	err := component.DecodeJSONBody(w, r, &book)
	if err != nil {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	fmt.Println(book.ID)
	fmt.Println(book.Title)
	json.NewEncoder(w).Encode(book)
	fmt.Println("ok")
}