package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"github.com/gorilla/mux"
	"selfregi/route"
	"selfregi/infrastructure/database"
)

func main() {

	r := mux.NewRouter()
	r = route.GetRoutes(r)
	db, _ := database.EntOpen()
	// database.Migrate(db)
	defer db.Close()
	
	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
  }
	fmt.Println("build ok")
	log.Fatal(server.ListenAndServe())
}