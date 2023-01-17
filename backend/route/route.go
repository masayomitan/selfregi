package route

import (
	// "net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	// "os"
	"fmt"
	FrontAPI "selfregi/controller/frontAPI"
	ManageAPI "selfregi/controller/manageAPI"
)

type Route func(*mux.Router)

func GetRoutes(r *mux.Router) *mux.Router{
	fmt.Println()
	r.Schemes("https")
	r.Headers("X-Requested-With", "XMLHttpRequest")

	err := godotenv.Load((".env"))
	if err != nil {
		panic ("envファイルの読み込みに失敗しました。")
	}

	// env := os.Getenv("ENV")
	r.HandleFunc("/top", FrontAPI.Top).Methods("GET")
	r.HandleFunc("/visitor", FrontAPI.VisitorStore).Methods("POST")
	
	// api
	api := r.PathPrefix("/api").Subrouter()

		// items
		item := api.PathPrefix("/items").Subrouter()
		item.HandleFunc("/get/{item_id:[0-9]+}", FrontAPI.GetItem).Methods("GET")
		
		// categories
		c := api.PathPrefix("/categories").Subrouter()
		c.HandleFunc("/{category_id:[0-9]+}/items/get", FrontAPI.GetItems).Methods("GET")
		c.HandleFunc("/get", FrontAPI.GetCategories).Methods("GET")

		///////////////// admin /////////////////
		admin := api.PathPrefix("/admin").Subrouter()
		admin.HandleFunc("/regist", ManageAPI.AdminRegist).Methods("POST")
		admin.HandleFunc("/update", ManageAPI.AdminUpdate).Methods("PATCH")

			// admin/categories
			adminCategory := admin.PathPrefix("/categories").Subrouter()
			adminCategory.HandleFunc("/get", ManageAPI.GetCategories).Methods("GET")

			// admin/items
			adminItem := admin.PathPrefix("/items").Subrouter()
			adminItem.HandleFunc("/get", ManageAPI.GetItems).Methods("GET")
			adminItem.HandleFunc("/post", ManageAPI.PostItem).Methods("POST")

	
	return r
}

