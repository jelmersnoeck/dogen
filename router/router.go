package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jelmersnoeck/noscito/controllers"
)

func HandleRoutes() {
	router := mux.NewRouter()

	router.HandleFunc(
		"/documents/{template}",
		controllers.DocumentsShow,
	).Methods("POST")

	http.Handle("/", router)
}
