package router

import (
	"uploader/cmd/handler"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Init ...
func Init() *mux.Router {

	var router = mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("API")

	})
	
		router.HandleFunc("/upload", handler.UploadFile) 

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(" 404 Not found ")
	})

	return router
}
