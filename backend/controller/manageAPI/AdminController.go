package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"selfregi/model"
	"selfregi/usecase"

	// "github.com/gorilla/mux"
	"io/ioutil"
	// "github.com/gorilla/sessions"
)

type test struct {
	Name string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

var admin model.Admin

func AdminRegist (w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	b, _ := json.Marshal(reqBody)
	_ = json.Unmarshal([]byte(b), &admin)
	fmt.Println(&admin)
	// fmt.Println(&admin.Password)
	ctx := r.Context()
	usecase.AdminRegist(ctx, admin)
}

// 構造体の修正必要
func AdminUpdate (w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	b, _ := json.Marshal(reqBody)
	_ = json.Unmarshal([]byte(b), &admin)
	fmt.Println(b)
	ctx := r.Context()
	usecase.AdminUpdate(ctx, admin)
}