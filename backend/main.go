package main

import (
	"net/http"
	"fmt"
)

func main() {
	server := http.Server{
		Addr: ":8000", 
	}
	fmt.Println("!eee")
	// http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}