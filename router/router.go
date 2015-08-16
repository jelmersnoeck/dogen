package router

import (
	"github.com/gorilla/mux"
	"github.com/jelmersnoeck/noscito/controllers"
	"net/http"
)

func HandleRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.MainIndex)

	http.Handle("/", router)
}
