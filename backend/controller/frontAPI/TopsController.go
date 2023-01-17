package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	// "github.com/gorilla/mux"
)

func Top(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	fmt.Fprintf(w, "hemmmmmllo top")
}