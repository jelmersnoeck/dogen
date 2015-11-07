package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jelmersnoeck/dogen/web/controllers"
)

func HandleRoutes() {
	router := mux.NewRouter()

	router.HandleFunc(
		"/documents/{template}",
		controllers.DocumentsShow,
	).Methods("POST")

	http.Handle("/", router)
}
